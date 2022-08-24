package worker

import (
	context "context"
	"fmt"
	"log"
	"time"
	"worker/dao"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	USER_ACCOUNT     = "user"
	CONTRACT_ACCOUNT = "contract"
)

func (bl *Worker) Start(ctx context.Context) error {
	header := make(chan *types.Header)
	subs, err := bl.ethClient.SubscribeNewHead(ctx, header)
	if err != nil {
		log.Printf("failed on subscribe for block header: %v", err)
		return err
	}
	defer subs.Unsubscribe()

	for {
		select {
		case err = <-subs.Err():
			log.Printf("failed when listening for block header: %v", err)
			return err
		case newHeader := <-header:
			go func() {
				err := bl.processBlock(ctx, newHeader)
				if err != nil {
					log.Printf("failed on process block: %v", err)
					panic(err)
				}
			}()
		}
	}
}

func (bl *Worker) processBlock(ctx context.Context, header *types.Header) error {
	log.Println("start processing block: ", header.Number.String())
	block, err := bl.ethClient.BlockByNumber(ctx, header.Number)

	// In experience, there will be some momment that can not find the block.
	// Retry getting the block is require.
	for retry := 1; err != nil; retry++ {
		time.Sleep(1 * time.Second)
		log.Printf("Try get block number #%d: %s", retry, header.Number.String())
		block, err = bl.ethClient.BlockByNumber(ctx, header.Number)
	}

	if err != nil {
		return fmt.Errorf("failed on get block number: %s\n%v", header.Number.String(), err)
	}

	// recursive slice is used for append `to` address
	// so that I can build the infinite graph from specific address.
	recursive := []string{}

	for i, tx := range block.Transactions() {
		from, err := bl.ethClient.TransactionSender(ctx, tx, block.Hash(), uint(i))
		if err != nil {
			return fmt.Errorf("failed on get sender: %v", err)
		}
		to := tx.To()

		// check if tracking `from`
		trackedFrom, err := bl.Tracking(ctx, from.Hex())
		if err != nil {
			return fmt.Errorf("failed on check if tracking `from` %s : %v", from.Hex(), err)
		}

		// check if tracking `to`
		trackedTo := false
		if to != nil {
			trackedTo, err = bl.Tracking(ctx, to.Hex())
			if err != nil {
				return fmt.Errorf("failed on check if tracking `to` %s : %v", to.Hex(), err)
			}
		}

		var receipt *types.Receipt
		if trackedFrom || trackedTo {
			// @TODO: optimize needed HERE!!
			receipt, err = bind.WaitMined(ctx, bl.ethClient, tx)
			if err != nil {
				return fmt.Errorf("failed on get receipt: %v", err)
			}

			if trackedFrom {
				err := bl.processTx(ctx, &from, to, tx, header, receipt, dao.OutEdge)
				if err != nil {
					return fmt.Errorf("failed on process OutEdge %v", err)
				}
				if to != nil {
					recursive = append(recursive, to.Hex())
				}
			}

			// process `to`
			if to != nil && trackedTo {
				err := bl.processTx(ctx, to, &from, tx, header, receipt, dao.InEdge)
				if err != nil {
					return fmt.Errorf("failed on process InEdge %v", err)
				}
			}
		}
	}

	err = bl.addRecursiveAddress(ctx, recursive)
	if err != nil {
		return fmt.Errorf("failed on add recursive address: %v", err)
	}
	log.Println("done processed block: ", header.Number.String())
	return nil
}

func (bl *Worker) addRecursiveAddress(ctx context.Context, recursive []string) error {
	for _, recAddr := range recursive {
		accountType, err := bl.GetAccountType(recAddr)
		if err != nil {
			return fmt.Errorf("failed on add recursive address (get account type method): %v", err)
		}
		if _, err := bl.InsertTrackedAddress(ctx, recAddr, accountType); err != nil {
			return fmt.Errorf("failed on add recursive address (insert method): %v", err)
		}
	}
	return nil
}

func (bl *Worker) processTx(ctx context.Context, from, to *common.Address, tx *types.Transaction, header *types.Header, receipt *types.Receipt, edgeDirect int) error {
	accountType := "contract"
	if from != nil {
		at, err := bl.GetAccountType(from.Hex())
		if err != nil {
			return err
		}
		accountType = at
	}
	txEdge, event := processTxEdge(tx, header, receipt, to, edgeDirect)
	if err := bl.dao.UpdateDB(ctx, &dao.Vertex{Address: from.Hex(), Type: accountType}, txEdge, event); err != nil {
		return err
	}
	return nil
}

func processTxEdge(tx *types.Transaction, header *types.Header, receipt *types.Receipt, addr *common.Address, edgeDirect int) (*dao.TxEdge, []*dao.Event) {
	createTime := time.Unix(int64(header.Time), 0).UTC()
	address := "contract creation"
	if addr != nil {
		address = addr.Hex()
	}
	status := receipt.Status
	txHash := tx.Hash().Hex()
	blockNumber := header.Number.String()
	value := tx.Value().String()
	direct := edgeDirect
	event := processEvent(receipt)
	return &dao.TxEdge{
		CreatedAt:   createTime,
		Address:     address,
		Status:      status,
		TxHash:      txHash,
		BlockNumber: blockNumber,
		Value:       value,
		Direct:      direct,
	}, event
}

func processEvent(receipt *types.Receipt) []*dao.Event {
	events := []*dao.Event{}
	for _, log := range receipt.Logs {
		addr := log.Address.Hex()
		topics := []string{}
		for _, t := range log.Topics {
			topics = append(topics, t.Hex())
		}
		logData := string(log.Data)
		events = append(events, &dao.Event{
			Address: addr,
			Topics:  topics,
			Data:    logData,
		})
	}
	return events
}
