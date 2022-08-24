package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"worker/dao"
	worker "worker/worker"

	"google.golang.org/grpc"
)

var (
	gRPCPort = os.Getenv("GRPC_PORT")
)

func main() {
	grpcServer := grpc.NewServer()

	ethConn, err := connectToEthereumNode(os.Getenv("NODE_URL"))
	if err != nil {
		panic(fmt.Errorf("failed on connect to Ethereum node: %v", err))
	}

	mongoConn, err := connectToMongoDB(os.Getenv("MONGODB_URI"), os.Getenv("MONGO_DATABASE"), os.Getenv("MONGO_USERNAME"), os.Getenv("MONGO_PASSWORD"))
	if err != nil {
		panic(fmt.Errorf("failed on connect to mongo: %v", err))
	}

	dao := dao.New(os.Getenv("MONGO_DATABASE"), connectToDB(mongoConn, os.Getenv("MONGO_DATABASE")))
	if err != nil {
		panic(fmt.Errorf("failed on get graphDB instance: %v", err))
	}

	blockWorker := worker.New(ethConn, mongoConn, dao)
	defer blockWorker.Close()

	worker.RegisterWorkerServer(grpcServer, blockWorker)

	log.Println("Start listening new block")
	go func() {
		err := blockWorker.Start(context.Background())
		if err != nil {
			panic(err)
		}
	}()

	defer grpcServer.Stop()

	// because run each server in goroutine so use a block channel to
	// prevent main thread cancel all its child.
	running := make(chan int)

	startServer := func(lis net.Listener, port string) {
		log.Printf("gRPC server is listening on port: %s", port)

		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("failed to listen on port %s: %v", port, err)
			panic(err)
		}
	}

	for _, port := range []string{"50001"} {
		listen, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
		if err != nil {
			log.Printf("failed to listen on port %s: %v", port, err)
			panic(err)
		}
		go startServer(listen, port)
	}

	<-running
}
