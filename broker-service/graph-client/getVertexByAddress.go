package graph

import (
	"context"
	"fmt"
)

func (graph *GraphClient) GetVertexByAddress(ctx context.Context, vertexRequest *VertexRequest) (*GinResponse, error) {
	resp, err := graph.client.GetVertexByAddress(ctx, vertexRequest)
	if err != nil {
		return createResponse(true, "FAIL", err.Error(), nil), err
	}
	return createResponse(false, "OK", fmt.Sprintf("sucessfully get address %s", vertexRequest.GetAddress()), *resp), nil
}
