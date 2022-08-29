package graph

import (
	"context"
	"fmt"
)

func (graph *GraphClient) GetTxByTxHash(ctx context.Context, query *Query) (*GinResponse, error) {
	resp, err := graph.client.GetTxByTxHash(ctx, query)
	if err != nil {
		return createResponse(true, "FAIL", err.Error(), nil), err
	}
	return createResponse(false, "OK", fmt.Sprintf("sucessfully get tx by TxHash: %s", query.GetTxHash()), *resp), nil
}
