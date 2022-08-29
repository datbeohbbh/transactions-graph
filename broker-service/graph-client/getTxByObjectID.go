package graph

import (
	"context"
	"fmt"
)

func (graph *GraphClient) GetTxByObjectID(ctx context.Context, query *Query) (*GinResponse, error) {
	resp, err := graph.client.GetTxByObjectID(ctx, query)
	if err != nil {
		return createResponse(true, "FAIL", err.Error(), nil), err
	}
	return createResponse(false, "OK", fmt.Sprintf("sucessfully get tx by objectID: %s", query.GetObjectID()), *resp), nil
}
