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
	case "get-by-account-type":
		b := graph.HandleGetAddressByAccountType(ctx, data)
		return b
	case "get-tx-by-object-id":
		b := graph.HandleGetTxByObjectID(ctx, data)
		return b
	case "get-tx-by-tx-hash":
		b := graph.HandleGetTxByTxHash(ctx, data)
		return b
	case "get-tx-by-address":
		b := graph.HandleGetTxByAddress(ctx, data)
		return b
	case "get-tx-by-block-number":
		b := graph.HandleGetTxByBlockNumber(ctx, data)
		return b
	case "get-tx-by-edge-direction":
		b := graph.HandleGetTxByEdgeDirection(ctx, data)
		return b
	case "get-tx-by-filter":
		b := graph.HandleGetTxByFilter(ctx, data)
		return b
	default:
		return readJson(*createResponse(true, "NOT SUPPORTED ACTION", fmt.Sprintf("action %s is not supported", action), nil))
	}
}

func (graph *GraphClient) HandleGetByAddress(ctx context.Context, data any) []byte {
	vertexRequest := Query{}
	b := readJson(data)
	writeJson(b, &vertexRequest)
	ginResp, _ := graph.GetVertexByAddress(ctx, &vertexRequest)
	b = readJson(*ginResp)
	return b
}

func (graph *GraphClient) HandleGetAddressByAccountType(ctx context.Context, data any) []byte {
	query := Query{}
	b := readJson(data)
	writeJson(b, &query)
	ginResp, _ := graph.GetVertexByAccountType(ctx, &query)
	b = readJson(*ginResp)
	return b
}

func (graph *GraphClient) HandleGetAllAddress(ctx context.Context, data any) []byte {
	ginResp, _ := graph.GetAllVertex(ctx, &Empty{})
	b := readJson(*ginResp)
	return b
}

func (graph *GraphClient) HandleGetTxByObjectID(ctx context.Context, data any) []byte {
	query := Query{}
	b := readJson(data)
	writeJson(b, &query)
	ginResp, _ := graph.GetTxByObjectID(ctx, &query)
	b = readJson(*ginResp)
	return b
}

func (graph *GraphClient) HandleGetTxByTxHash(ctx context.Context, data any) []byte {
	query := Query{}
	b := readJson(data)
	writeJson(b, &query)
	ginResp, _ := graph.GetTxByTxHash(ctx, &query)
	b = readJson(*ginResp)
	return b
}

func (graph *GraphClient) HandleGetTxByAddress(ctx context.Context, data any) []byte {
	query := Query{}
	b := readJson(data)
	writeJson(b, &query)
	ginResp, _ := graph.GetTxByAddress(ctx, &query)
	b = readJson(*ginResp)
	return b
}

func (graph *GraphClient) HandleGetTxByBlockNumber(ctx context.Context, data any) []byte {
	query := Query{}
	b := readJson(data)
	writeJson(b, &query)
	ginResp, _ := graph.GetTxByBlockNumber(ctx, &query)
	b = readJson(*ginResp)
	return b
}

func (graph *GraphClient) HandleGetTxByEdgeDirection(ctx context.Context, data any) []byte {
	query := Query{}
	b := readJson(data)
	writeJson(b, &query)
	ginResp, _ := graph.GetTxByEdgeDirection(ctx, &query)
	b = readJson(*ginResp)
	return b
}

func (graph *GraphClient) HandleGetTxByFilter(ctx context.Context, data any) []byte {
	filters := Filters{}
	b := readJson(data)
	writeJson(b, &filters)
	ginResp, _ := graph.GetTxByFilter(ctx, &filters)
	b = readJson(*ginResp)
	return b
}
