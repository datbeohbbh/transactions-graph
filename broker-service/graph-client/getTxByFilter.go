package graph

import (
	"context"
	"fmt"
)

func (graph *GraphClient) GetTxByFilter(ctx context.Context, filters *Filters) (*GinResponse, error) {
	resp, err := graph.client.GetTxByFilter(ctx, filters)
	if err != nil {
		return createResponse(true, "FAIL", err.Error(), nil), err
	}
	return createResponse(false, "OK", fmt.Sprintf("sucessfully get tx by filter: %+v", filters), *resp), nil
}
