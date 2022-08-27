package main

import (
	"broker/address-manager"
	"broker/graph-client"
	"context"
	"fmt"
)

func (broker *Broker) HandleRequest(ctx context.Context, request *Request) *Response {
	switch request.Class {
	case "address":
		return broker.AddressManagerHandler(broker.addressManagerClient, ctx, request.Action, request.Data)
	case "graph":
		return broker.GraphHandler(broker.graphClient, ctx, request.Action, request.Data)
	default:
		return createResponse(true, "NOT SUPPORTED CLASS", fmt.Sprintf("class %s is not supported", request.Class), nil)
	}
}

func byteToResponse(b []byte) *Response {
	resp := Response{}
	writeJson(b, &resp)
	return &resp
}

func (broker *Broker) AddressManagerHandler(client *address.Client, ctx context.Context, action string, data any) *Response {
	b := client.Handle(ctx, action, data)
	return byteToResponse(b)
}

func (broker *Broker) GraphHandler(client *graph.GraphClient, ctx context.Context, action string, data any) *Response {
	b := client.Handle(ctx, action, data)
	return byteToResponse(b)
}
