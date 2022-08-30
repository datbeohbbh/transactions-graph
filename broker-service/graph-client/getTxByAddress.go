package graph

import (
	"context"
	"fmt"
)

func (graph *GraphClient) GetTxByAddress(ctx context.Context, query *Query) (*GinResponse, error) {
	resp, err := graph.client.GetTxByAddress(ctx, query)
	if err != nil {
		return createResponse(true, "FAIL", err.Error(), nil), err
	}
	return createResponse(false, "OK", fmt.Sprintf("sucessfully get tx by address: %s", query.GetAddress()), resp), nil
}
