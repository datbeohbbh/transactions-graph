package main

import (
	"context"
	"fmt"
)

func (broker *Broker) HandleRequest(ctx context.Context, request *Request) *Response {
	switch request.Class {
	case "address":
		b := broker.addressManagerClient.Handle(ctx, request.Action, request.Data)
		resp := Response{}
		writeJson(b, &resp)
		return &resp
	case "graph":
		return nil
		/* 		b, err := broker.graphClient.Handle(ctx, request.Action, request.Data)
		   		if err != nil {
		   			return nil, err
		   		}
		   		resp := Response{}
		   		writeJson(b, resp)
		   		return &resp, nil */
	default:
		return createResponse(true, "NOT SUPPORTED CLASS", fmt.Sprintf("class %s is not supported", request.Class), nil)
	}
}
