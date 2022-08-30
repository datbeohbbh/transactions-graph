package worker

import (
	"context"
	"github.com/datbeohbbh/transactions-graph/worker/dao"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (worker *Worker) processTx(ctx context.Context, from, to *common.Address, tx *types.Transaction, header *types.Header, receipt *types.Receipt, edgeDirect int) error {
	accountType := "contract"
	if from != nil {
		at, err := worker.addressClient.AccountType(ctx, from.Hex())
		if err != nil {
			return err
		}
		accountType = at
	}
	txEdge, event := processTxEdge(tx, header, receipt, to, edgeDirect)
	if err := worker.dao.UpdateDB(ctx, &dao.Vertex{Address: from.Hex(), Type: accountType}, txEdge, event); err != nil {
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
