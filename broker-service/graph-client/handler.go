package graph

import (
	context "context"
	"fmt"
)

type GinResponse struct {
	Error   bool   `json:"error"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func createResponse(Error_ bool, Status_, Message_ string, Data_ any) *GinResponse {
	return &GinResponse{
		Error:   Error_,
		Status:  Status_,
		Message: Message_,
		Data:    Data_,
	}
}

func (graph *GraphClient) Handle(ctx context.Context, action string, data any) []byte {
	switch action {
	case "get-by-address":
		b := graph.HandleGetByAddress(ctx, data)
		return b
	case "get-all-address":
		b := graph.HandleGetAllAddress(ctx, data)
		return b
	default:
		return readJson(*createResponse(true, "NOT SUPPORTED ACTION", fmt.Sprintf("action %s is not supported", action), nil))
	}
}

func (graph *GraphClient) HandleGetByAddress(ctx context.Context, data any) []byte {
	vertexRequest := VertexRequest{}
	b := readJson(data)
	writeJson(b, &vertexRequest)
	ginResp, _ := graph.GetVertexByAddress(ctx, &vertexRequest)
	b = readJson(*ginResp)
	return b
}

func (graph *GraphClient) HandleGetAllAddress(ctx context.Context, data any) []byte {
	ginResp, _ := graph.GetAllVertex(ctx, &Empty{})
	b := readJson(*ginResp)
	return b
}
