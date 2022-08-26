package address

import (
	"context"
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

func (addressClient *Client) Handle(ctx context.Context, action string, data any) []byte {
	switch action {
	case "add-address":
		b := addressClient.HandleAddAddress(ctx, data)
		return b
	case "remove-address":
		b := addressClient.HandleRemoveAddress(ctx, data)
		return b
	case "check-tracking":
		b := addressClient.HandleIsTrackingAddress(ctx, data)
		return b
	default:
		return readJson(*createResponse(true, "NOT SUPPORTED ACTION", fmt.Sprintf("action %s is not supported", action), nil))
	}
}

func (addressClient *Client) HandleAddAddress(ctx context.Context, data any) []byte {
	addr := Address{}
	b := readJson(data)
	writeJson(b, &addr)
	ginResp, _ := addressClient.AddAddress(ctx, &addr)
	b = readJson(*ginResp)
	return b
}

func (addressClient *Client) HandleRemoveAddress(ctx context.Context, data any) []byte {
	addr := Address{}
	b := readJson(data)
	writeJson(b, &addr)
	ginResp, _ := addressClient.RemoveAddress(ctx, &addr)
	b = readJson(*ginResp)
	return b
}

func (addressClient *Client) HandleIsTrackingAddress(ctx context.Context, data any) []byte {
	addr := Address{}
	b := readJson(data)
	writeJson(b, &addr)
	ginResp, _ := addressClient.IsTracking(ctx, &addr)
	b = readJson(*ginResp)
	return b
}
