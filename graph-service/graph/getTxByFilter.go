package graph

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (graph *GraphData) GetTxByFilter(ctx context.Context, filters *Filters) (*Txs, error) {
	filterList := filters.GetFilter()
	conds := make(map[string]string)
	for _, f := range filterList {
		conds[f.GetKey()] = f.GetValue()
	}

	txs, err := graph.dao.GetTxByFilter(ctx, conds)
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
			Direct:      Tx_Direction_name[int32(tx.Direct)],
		})
	}
	return &result, nil
}
