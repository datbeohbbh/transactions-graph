package worker

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/datbeohbbh/transactions-graph/worker/dao"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

func (worker *Worker) processBlock(ctx context.Context, header *types.Header) error {
	log.Println("start processing block: ", header.Number.String())
	block, err := worker.ethClient.BlockByNumber(ctx, header.Number)

	// In experience, there will be some momment that can not find the block.
	// Retry getting the block is require.
	for retry := 1; err != nil; retry++ {
		time.Sleep(1 * time.Second)
		log.Printf("Try get block number #%d: %s", retry, header.Number.String())
		block, err = worker.ethClient.BlockByNumber(ctx, header.Number)
	}

	if err != nil {
		return fmt.Errorf("failed on get block number: %s\n%v", header.Number.String(), err)
	}

	// recursive slice is used for append `to` address
	// so that I can build the infinite graph from specific address.
	recursive := []string{}

	for i, tx := range block.Transactions() {
		from, err := worker.ethClient.TransactionSender(ctx, tx, block.Hash(), uint(i))
		if err != nil {
			return fmt.Errorf("failed on get sender: %v", err)
		}
		to := tx.To()

		// check if tracking `from`
		trackedFrom, err := worker.addressClient.IsTracking(ctx, from.Hex())
		if err != nil {
			return fmt.Errorf("failed on check if tracking `from` %s : %v", from.Hex(), err)
		}

		// check if tracking `to`
		trackedTo := false
		if to != nil {
			trackedTo, err = worker.addressClient.IsTracking(ctx, to.Hex())
			if err != nil {
				return fmt.Errorf("failed on check if tracking `to` %s : %v", to.Hex(), err)
			}
		}

		var receipt *types.Receipt
		if trackedFrom || trackedTo {
			// @TODO: optimize needed HERE!!
			receipt, err = bind.WaitMined(ctx, worker.ethClient, tx)
			if err != nil {
				return fmt.Errorf("failed on get receipt: %v", err)
			}

			if trackedFrom {
				// in TxEdge: from -> to, to mark as IN edge bc it is a directed edge from `from` to `to`.
				err := worker.processTx(ctx, &from, to, tx, header, receipt, dao.IN)
				if err != nil {
					return fmt.Errorf("failed on process OutEdge %v", err)
				}
				if to != nil {
					recursive = append(recursive, to.Hex())
				}
			}

			// process `to`
			if to != nil && trackedTo {
				err := worker.processTx(ctx, to, &from, tx, header, receipt, dao.OUT)
				if err != nil {
					return fmt.Errorf("failed on process InEdge %v", err)
				}
			}
		}
	}

	countRecursiveAdd, err := worker.addRecursiveAddress(ctx, recursive)
	if err != nil {
		return fmt.Errorf("failed on add recursive address: %v", err)
	}
	log.Printf("block #%s: added %d recursive address", header.Number.String(), countRecursiveAdd)
	return nil
}

func (worker *Worker) addRecursiveAddress(ctx context.Context, recursive []string) (int, error) {
	ret := 0
	for _, recAddr := range recursive {
		if f, err := worker.addressClient.AddAddress(ctx, recAddr); err != nil {
			return 0, fmt.Errorf("failed on add recursive address (insert method): %v", err)
		} else if f {
			ret += 1
		}
	}
	return ret, nil
}
