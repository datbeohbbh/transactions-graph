package graph

import (
	"context"
	"fmt"
)

func (graph *GraphClient) GetTxByEdgeDirection(ctx context.Context, query *Query) (*GinResponse, error) {
	resp, err := graph.client.GetTxByEdgeDirection(ctx, query)
	if err != nil {
		return createResponse(true, "FAIL", err.Error(), nil), err
	}
	return createResponse(false, "OK",
		fmt.Sprintf("sucessfully get tx by edge direction: %s (0 - IN / 1 - OUT)",
			query.GetDirect()),
		resp), nil
}
