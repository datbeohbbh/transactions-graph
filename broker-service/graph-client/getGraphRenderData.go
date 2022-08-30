package graph

import (
	"context"

	"github.com/datbeohbbh/transactions-graph/broker/graphAlgo"
)

func (graph *GraphClient) GetGraphRenderData(ctx context.Context, query *graphAlgo.Query) (*GinResponse, error) {
	result, err := graph.client.GetGraphRenderData(ctx, query)
	if err != nil {
		return createResponse(true, "FAIL", err.Error(), nil), err
	}
	return createResponse(false, "OK", "sucessfully build graph", result), nil
}
