package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	addressManager "github.com/datbeohbbh/transactions-graph/worker/address-manager"
	"github.com/datbeohbbh/transactions-graph/worker/consumers"
	"github.com/datbeohbbh/transactions-graph/worker/dao"
	worker "github.com/datbeohbbh/transactions-graph/worker/worker"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	defer ethConn.Close()

	mongoConn, err := connectToMongoDB(
		os.Getenv("MONGODB_URI"),
		os.Getenv("MONGO_DATABASE"),
		os.Getenv("MONGO_USERNAME"),
		os.Getenv("MONGO_PASSWORD"))
	if err != nil {
		panic(fmt.Errorf("failed on connect to mongo: %v", err))
	}
	defer mongoConn.Disconnect(context.Background())

	dao := dao.New(os.Getenv("MONGO_DATABASE"), connectToDB(mongoConn, os.Getenv("MONGO_DATABASE")))
	if err != nil {
		panic(fmt.Errorf("failed on get graphDB instance: %v", err))
	}

	addressManagerRPCConn, err := connectToGRPCServerContext(
		context.Background(),
		os.Getenv("NGINX_SERVICE_NAME"),
		os.Getenv("NGINX_GRPC_PORT"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		panic(fmt.Errorf("failed on connect to address manager RPC server: %v", err))
	}

	addressClient := addressManager.New(addressManagerRPCConn)

	amqpConn, err := connectToRabbitMQ(
		os.Getenv("AMQP_USERNAME"),
		os.Getenv("AMQP_PASSWORD"),
		os.Getenv("AMQP_HOST"),
		os.Getenv("AMQP_PORT"))
	if err != nil {
		panic(fmt.Errorf("failed on connect to rabbitmq: %v", err))
	}
	defer amqpConn.Close()

	consumer, err := consumers.New(amqpConn)
	if err != nil {
		panic(fmt.Errorf("failed on create consumer: %v", err))
	}
	consumer.SetUp()
	defer consumer.Close()

	blockWorker := worker.New(ethConn, dao, addressClient, consumer)

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
