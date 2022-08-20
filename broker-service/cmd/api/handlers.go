package main

import (
	pb "broker/listener"
	"context"
	"fmt"
)

func (broker *Broker) HandleRequest(ctx context.Context, action, address string) (*Response, error) {
	switch action {
	case "add-address":
		resp, err := broker.HandleAddAddress(ctx, address)
		return resp, err
	case "graph":
		return nil, nil
	default:
		return nil, fmt.Errorf("action %s is not supported", action)
	}
}

func (broker *Broker) HandleAddAddress(ctx context.Context, address string) (*Response, error) {
	client := pb.NewBlockListenerClient(broker.grpcConn)
	resp, err := client.AddAddress(ctx, &pb.Address{Address: address})
	if err != nil {
		return nil, err
	}
	payload := Response{
		Status:  resp.Status,
		Message: resp.Msg,
	}
	return &payload, nil
}
