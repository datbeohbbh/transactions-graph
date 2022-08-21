package main

import (
	"block-listener/dao"
	"block-listener/listener"
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

var (
	gRPCPort = os.Getenv("GRPC_PORT")
)

func main() {
	if gRPCPort == "" {
		panic(errors.New("undefined gRPC port"))
	}
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", gRPCPort))
	if err != nil {
		log.Printf("failed to listen on port %s: %v", gRPCPort, err)
		panic(err)
	}

	grpcServer := grpc.NewServer()

	ethConn, err := connectToEthereumNode(os.Getenv("NODE_URL"))
	if err != nil {
		panic(fmt.Errorf("failed on connect to Ethereum node: %v", err))
	}

	mongoConn, err := connectToMongoDB(os.Getenv("MONGODB_URI"), os.Getenv("MONGO_DATABASE"), os.Getenv("MONGO_USERNAME"), os.Getenv("MONGO_PASSWORD"))
	if err != nil {
		panic(fmt.Errorf("failed on connect to mongo: %v", err))
	}

	graphDBInstance, err := connectToGraphDB(mongoConn, os.Getenv("MONGO_DATABASE"))
	if err != nil {
		panic(fmt.Errorf("failed on get graphDB instance: %v", err))
	}

	blockListenerHandler := listener.New(ethConn, mongoConn, graphDBInstance)
	defer blockListenerHandler.Close()

	mongoConn, err = connectToMongoDB(os.Getenv("DAO_MONGODB_URI"), os.Getenv("DAO_MONGO_DATABASE"), os.Getenv("DAO_MONGO_USERNAME"), os.Getenv("DAO_MONGO_PASSWORD"))
	if err != nil {
		panic(fmt.Errorf("failed on connect to mongo db for DAO"))
	}
	daoHandler := dao.New(mongoConn, graphDBInstance)
	defer daoHandler.Close()

	listener.RegisterBlockListenerServer(grpcServer, blockListenerHandler)
	dao.RegisterDAOServer(grpcServer, daoHandler)

	log.Printf("gRPC server is listening on port: %s", gRPCPort)
	defer grpcServer.Stop()

	LOW, err := scaleMethod(context.Background())
	if err != nil || LOW == -1 {
		panic(fmt.Errorf("failon on get range for scaling: %d - %v", LOW, err))
	}
	log.Printf("serving block with last digit in range [%d, %d): ", 2*LOW, 2*(LOW+1))

	log.Println("Start listening new block")
	go func() {
		err := blockListenerHandler.Start(context.Background(), 2*LOW, 2*(LOW+1))
		if err != nil {
			panic(err)
		}
	}()

	if err := grpcServer.Serve(listen); err != nil {
		log.Printf("failed to listen on port %s: %v", gRPCPort, err)
		panic(err)
	}
}
