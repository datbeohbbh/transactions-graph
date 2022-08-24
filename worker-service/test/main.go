package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	pb "block-listener/graph-data"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	gRPCPort = 50052
)

func main() {
	// must provide credentials to connect otherwise connect will fail.
	conn, err := grpc.Dial(fmt.Sprintf("0.0.0.0:%d", gRPCPort), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatal("failed to connect")
	}
	defer conn.Close()

	cl := pb.NewGraphDataClient(conn)
	for i, v := range []string{"0x28C6c06298d514Db089934071355E5743bf21d60", "0x65452920B76dd5862026a19582A77f95fB9a0E27", "0xcA302dE5463F562aD83ba59DD1e7eEB60ffC7142"} {
		resp, err := cl.GetVertexByAddress(context.Background(), &pb.VertexRequest{Address: v})
		if err != nil {
			log.Printf("fail to add address: %v", err)
		} else {
			s, _ := json.MarshalIndent(resp, "", "\t")
			log.Printf("test: #%d - msg: %s", i, string(s))
		}
	}
}
