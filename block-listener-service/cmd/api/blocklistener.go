package main

import (
	"block-listener/data"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (bl *BlockListener) blockListener(ctx context.Context) error {
	header := make(chan *types.Header)
	subs, err := bl.ethClient.SubscribeNewHead(ctx, header)
	if err != nil {
		log.Printf("failed on subscribe for block header: %v", err)
		return err
	}

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

func (bl *BlockListener) processBlock(ctx context.Context, header *types.Header) error {
	log.Println("start process block: ", header.Number.String())
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
		receipt, err := bind.WaitMined(ctx, bl.ethClient, tx)
		if err != nil {
			return fmt.Errorf("failed on get receipt: %v", err)
		}

		to := tx.To()
		from, err := bl.ethClient.TransactionSender(ctx, tx, block.Hash(), uint(i))
		if err != nil {
			return fmt.Errorf("failed on get sender: %v", err)
		}

		if has, _ := bl.levdb.Has([]byte(from.Hex())); has {
			//log.Printf("Block: %s - update out edge - %s", header.Number.String(), from.Hex())
			err := bl.processTx(ctx, &from, to, tx, header, receipt, data.OutEdge)
			if err != nil {
				return fmt.Errorf("failed on process OutEdge %v", err)
			}
			if to != nil {
				recursive = append(recursive, to.Hex())
			}
		}
		if to != nil {
			if has, _ := bl.levdb.Has([]byte(to.Hex())); has {
				//log.Printf("Block: %s - update in edge - %s", header.Number.String(), to.Hex())
				err := bl.processTx(ctx, to, &from, tx, header, receipt, data.InEdge)
				if err != nil {
					return fmt.Errorf("failed on process InEdge %v", err)
				}
			}
		}
	}

	err = bl.addRecursiveAddress(recursive)
	if err != nil {
		return fmt.Errorf("failed on add recursive address: %v", err)
	}
	log.Println("done process block: ", header.Number.String())

	return nil
}

func (bl *BlockListener) addRecursiveAddress(recursive []string) error {
	for _, recAddr := range recursive {
		if has, err := bl.HasKey(recAddr); err != nil {
			return fmt.Errorf("failed on add recursive address (has method): %v", err)
		} else if has {
			continue
		}
		accountType, err := bl.GetAccountType(recAddr)
		if err != nil {
			return fmt.Errorf("failed on add recursive address (get account type method): %v", err)
		}
		if err := bl.PutKey(recAddr, accountType); err != nil {
			return fmt.Errorf("failed on add recursive address (put method): %v", err)
		}
	}
	return nil
}

func (bl *BlockListener) processTx(ctx context.Context, from, to *common.Address, tx *types.Transaction, header *types.Header, receipt *types.Receipt, edgeDirect int) error {
	accountType := "contract creation"
	if from != nil {
		at, err := bl.GetAccountType(from.Hex())
		if err != nil {
			return err
		}
		accountType = at
	}
	txEdge, event := processTxEdge(tx, header, receipt, to, edgeDirect)
	if err := bl.graphDB.UpdateDB(ctx, &data.Vertex{Address: from.Hex(), Type: accountType}, txEdge, event); err != nil {
		return err
	}
	return nil
}

func processTxEdge(tx *types.Transaction, header *types.Header, receipt *types.Receipt, addr *common.Address, edgeDirect int) (*data.TxEdge, []*data.Event) {
	createTime := time.Unix(int64(header.Time), 0).UTC()
	address := ""
	if addr != nil {
		address = addr.Hex()
	}
	status := receipt.Status
	txHash := tx.Hash().Hex()
	blockNumber := header.Number.String()
	value := tx.Value().String()
	direct := edgeDirect
	event := processEvent(receipt)
	return &data.TxEdge{
		CreatedAt:   createTime,
		Address:     address,
		Status:      status,
		TxHash:      txHash,
		BlockNumber: blockNumber,
		Value:       value,
		Direct:      direct,
	}, event
}

func processEvent(receipt *types.Receipt) []*data.Event {
	events := []*data.Event{}
	for _, log := range receipt.Logs {
		addr := log.Address.Hex()
		topics := []string{}
		for _, t := range log.Topics {
			topics = append(topics, t.Hex())
		}
		logData := string(log.Data)
		events = append(events, &data.Event{
			Address: addr,
			Topics:  topics,
			Data:    logData,
		})
	}
	return events
}
