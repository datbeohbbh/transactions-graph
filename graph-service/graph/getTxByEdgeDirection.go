package graph

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (graph *GraphData) GetTxByEdgeDirection(ctx context.Context, query *Query) (*Txs, error) {
	d, ok := Tx_Direction_value[query.GetDirect()]
	if !ok {
		return nil, fmt.Errorf("failed on get direct: %s (direction should be IN/OUT)", query.GetDirect())
	}
	txs, err := graph.dao.GetTxByEdgeDirection(ctx, d)
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
