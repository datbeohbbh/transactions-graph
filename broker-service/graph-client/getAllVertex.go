package graph

import "context"

func (graph *GraphClient) GetAllVertex(ctx context.Context, emptyField *Empty) (*GinResponse, error) {
	resp, err := graph.client.GetAllVertex(ctx, emptyField)
	if err != nil {
		return createResponse(true, "FAIL", err.Error(), nil), err
	}
	return createResponse(false, "OK", "successfully get all vertex", *resp), nil
}
