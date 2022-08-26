package graph

import "google.golang.org/grpc"

type GraphClient struct {
	client GraphDataClient
}

func New(grpcConn *grpc.ClientConn) *GraphClient {
	return &GraphClient{client: NewGraphDataClient(grpcConn)}
}
