package graph

import (
	"context"
	"fmt"
)

func (graph *GraphClient) GetVertexByAccountType(ctx context.Context, query *Query) (*GinResponse, error) {
	resp, err := graph.client.GetVertexByAccountType(ctx, query)
	if err != nil {
		return createResponse(true, "FAIL", err.Error(), nil), err
	}
	return createResponse(false, "OK", fmt.Sprintf("sucessfully get address by type: %s", query.GetAccountType()), resp), nil
}
