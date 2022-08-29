package graph

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (graph *GraphData) GetTxByAddress(ctx context.Context, query *Query) (*Txs, error) {
	txs, err := graph.dao.GetTxByAddress(ctx, query.GetAddress())
	if err != nil {
		return nil, err
	}

	result := Txs{}
	for _, tx := range txs {
		result.Txs = append(result.Txs, &Tx{
			CreatedAt:   timestamppb.New(tx.CreatedAt),
			Address:     tx.Address,
			Status:      int64(tx.Status),
			TxHash:      tx.TxHash,
			BlockNumber: tx.BlockNumber,
			Value:       tx.Value,
			EventLog:    tx.EventLog,
			Direct:      int32(tx.Direct),
		})
	}
	return &result, nil
}
