package graph

import (
	"context"

	"github.com/datbeohbbh/transactions-graph/graph-service/graphAlgo"
)

func (graph *GraphData) GetGraphRenderData(ctx context.Context, query *graphAlgo.Query) (*graphAlgo.GraphRenderData, error) {
	graph.renderContext.SetGraphAlgo(graphAlgo.NewBfs(
		ctx,
		graph.dao,
		query.GetFrom(),
		query.GetDepth(),
		query.GetTxCompletedBefore().AsTime(),
	))

	result, err := graph.renderContext.Execute()
	if err != nil {
		return nil, err
	}
	return result, nil
}
