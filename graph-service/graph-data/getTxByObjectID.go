package graph

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (graph *GraphData) GetTxByObjectID(ctx context.Context, query *Query) (*Tx, error) {
	tx, err := graph.dao.GetTxByObjectID(ctx, query.GetObjectID())
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
