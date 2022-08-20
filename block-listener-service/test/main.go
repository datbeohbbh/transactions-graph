package main

import (
	"context"
	"fmt"
	"log"

	pb "block-listener/listener"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	gRPCPort = 50001
)

func main() {
	// must provide credentials to connect otherwise connect will fail.
	conn, err := grpc.Dial(fmt.Sprintf(":%d", gRPCPort), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatal("failed to connect")
	}
	defer conn.Close()

	cl := pb.NewBlockListenerClient(conn)
	for i, v := range []string{"0x28C6c06298d514Db089934071355E5743bf21d60", "0x65452920B76dd5862026a19582A77f95fB9a0E27", "0xcA302dE5463F562aD83ba59DD1e7eEB60ffC7142"} {
		resp, err := cl.AddAddress(context.Background(), &pb.Address{Address: v})
		if err != nil {
			log.Printf("fail to add address: %v", err)
		} else {
			log.Printf("test: #%d - msg: %s", i, resp.Msg)
		}
	}
}
