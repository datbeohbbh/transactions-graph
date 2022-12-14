package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/datbeohbbh/transactions-graph/graph-service/dao"
	gr "github.com/datbeohbbh/transactions-graph/graph-service/graph"
)

var (
	GRPC_PORT = os.Getenv("GR_GRPC_PORT")
)

func main() {
	mongoConn, err := connectToMongoDB(os.Getenv("GR_MONGODB_URI"), os.Getenv("GR_MONGO_DATABASE"), os.Getenv("GR_MONGO_USERNAME"), os.Getenv("GR_MONGO_PASSWORD"))
	if err != nil {
		panic(fmt.Errorf("failed on connect to mongo: %v", err))
	}
	defer mongoConn.Disconnect(context.Background())

	dao := dao.New(os.Getenv("GR_MONGO_DATABASE"), connectToDB(mongoConn, os.Getenv("GR_MONGO_DATABASE")))
	if err != nil {
		panic(fmt.Errorf("failed on get graphDB instance: %v", err))
	}

	graphHandler := gr.New(dao)

	grpcServer := grpc.NewServer()

	gr.RegisterGraphDataServer(grpcServer, graphHandler)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", GRPC_PORT))
	if err != nil {
		panic(fmt.Errorf("failed on get listen tcp: %v", err))
	}

	log.Printf("Start listening on port: %s", GRPC_PORT)
	if err := grpcServer.Serve(lis); err != nil {
		panic(fmt.Errorf("failed on serve grpc server: %v", err))
	}
}
