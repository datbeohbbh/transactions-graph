package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	addressManager "github.com/datbeohbbh/transactions-graph/address-manger/address-manager"
	"github.com/datbeohbbh/transactions-graph/address-manger/dao"

	"google.golang.org/grpc"
)

var (
	GRPC_PORT = os.Getenv("ADDR_GRPC_PORT")

	DB_NAME = os.Getenv("DB_NAME")
)

func failedOnError(msg string, err error) {
	if err != nil {
		panic(fmt.Errorf("error: %s - %v", msg, err))
	}
}

func main() {
	mongoConn, err := connectToMongoDB(
		os.Getenv("ADDR_MONGODB_URI"),
		os.Getenv("ADDR_MONGO_DATABASE"),
		os.Getenv("ADDR_MONGO_USERNAME"),
		os.Getenv("ADDR_MONGO_PASSWORD"))

	failedOnError("failed on connect to mongoDB", err)
	defer mongoConn.Disconnect(context.Background())

	db := connectToDB(mongoConn, DB_NAME)
	failedOnError("failed on connect to graphDB", err)

	daoInstance := dao.New(DB_NAME, db)

	ethConn, err := connectToEthereumNode(os.Getenv("NODE_URL"))
	if err != nil {
		panic(fmt.Errorf("failed on connect to Ethereum node: %v", err))
	}
	defer ethConn.Close()

	handler := addressManager.New(daoInstance, ethConn)

	grpcServer := grpc.NewServer()

	addressManager.RegisterAddressManagerServer(grpcServer, handler)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", GRPC_PORT))
	failedOnError(fmt.Sprintf("failed on listen on port %s", GRPC_PORT), err)

	log.Printf("Start serving on port %s", GRPC_PORT)

	if err := grpcServer.Serve(lis); err != nil {
		failedOnError("faild on start grpc server", err)
	}
}
