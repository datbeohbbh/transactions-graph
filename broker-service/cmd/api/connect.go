package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

func connectToGRPCServerContext(ctx context.Context, host, port string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	connCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(connCtx, fmt.Sprintf("%s:%s", host, port), opts...)
	if err != nil {
		return nil, err
	}
	log.Printf("connected to gRPC server %s:%s", host, port)
	return conn, nil
}

func connectToGRPCServer(host, port string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), opts...)
	if err != nil {
		return nil, err
	}
	log.Printf("connected to gRPC server %s:%s", host, port)
	return conn, nil
}
