package graph

import (
	"context"
	"fmt"
)

func (graph *GraphClient) GetTxByBlockNumber(ctx context.Context, query *Query) (*GinResponse, error) {
	resp, err := graph.client.GetTxByBlockNumber(ctx, query)
	if err != nil {
		return createResponse(true, "FAIL", err.Error(), nil), err
	}
	return createResponse(false, "OK", fmt.Sprintf("sucessfully get tx by block number: %s", query.GetBlockNumber()), *resp), nil
}
