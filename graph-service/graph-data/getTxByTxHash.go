package graph

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (graph *GraphData) GetTxByTxHash(ctx context.Context, query *Query) (*Tx, error) {
	tx, err := graph.dao.GetTxByTxHash(ctx, query.GetTxHash())
	if err != nil {
		return nil, err
	}

	result := &Tx{
		CreatedAt:   timestamppb.New(tx.CreatedAt),
		Address:     tx.Address,
		Status:      int64(tx.Status),
		TxHash:      tx.TxHash,
		BlockNumber: tx.BlockNumber,
		Value:       tx.Value,
		EventLog:    tx.EventLog,
		Direct:      int32(tx.Direct),
	}

	return result, nil
}
