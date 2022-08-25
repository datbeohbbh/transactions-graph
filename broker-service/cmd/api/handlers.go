package main

import (
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
	status, msg, err := broker.addressManagerClient.AddAddress(ctx, address)
	if err != nil {
		return nil, err
	}
	return &Response{
		Error:   false,
		Status:  status,
		Message: msg,
	}, nil
}
