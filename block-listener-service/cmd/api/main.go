package main

import (
	"block-listener/data"
	pb "block-listener/listener"
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethdb/leveldb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var (
	gRPCPort     = os.Getenv("GRPC_PORT")
	GOERLI_HTTPS = "https://goerli.infura.io/v3/e8cc055ae4b5409ab1b43e9c0cefe56a"
	GOERLI_WSS   = "wss://goerli.infura.io/ws/v3/e8cc055ae4b5409ab1b43e9c0cefe56a"
)

const (
	USER_ACCOUNT     = "user"
	CONTRACT_ACCOUNT = "contract"
)

type BlockListener struct {
	pb.UnimplementedBlockListenerServer

	levdb *leveldb.Database

	ethClient *ethclient.Client

	mongoClient *mongo.Client

	graphDB *data.GraphDB
}

func gRPCHandler() (*BlockListener, error) {
	conn, err := connectToLevelDB()
	if err != nil {
		return nil, err
	}

	ethConn, err := connectToEthNode()
	if err != nil {
		return nil, err
	}

	mongoConn, err := connectToMongoDB()
	if err != nil {
		return nil, err
	}

	graphDBInstance, err := connectToGraphDB(mongoConn)
	if err != nil {
		return nil, err
	}

	return &BlockListener{
		levdb:       conn,
		ethClient:   ethConn,
		mongoClient: mongoConn,
		graphDB:     graphDBInstance,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", gRPCPort))
	if err != nil {
		log.Printf("failed to listen on port %s: %v", gRPCPort, err)
		panic(err)
	}

	grpcServer := grpc.NewServer()

	gRPCHandler, err := gRPCHandler()
	if err != nil {
		log.Printf("failed on create gRPC handler: %v", err)
		panic(err)
	}

	defer func() {
		if err := gRPCHandler.levdb.Close(); err != nil {
			panic(err)
		}
		gRPCHandler.ethClient.Close()
		if err := gRPCHandler.mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	pb.RegisterBlockListenerServer(grpcServer, gRPCHandler)

	log.Printf("gRPC server is listening on port: %s", gRPCPort)
	defer grpcServer.Stop()

	log.Println("Start listening new block")
	go func() {
		err := gRPCHandler.blockListener(context.Background())
		if err != nil {
			panic(err)
		}
	}()

	if err := grpcServer.Serve(listen); err != nil {
		log.Printf("failed to listen on port %s: %v", gRPCPort, err)
		panic(err)
	}
}

func connectToLevelDB() (*leveldb.Database, error) {
	db, err := leveldb.New("./db-data/leveldb", 0, 0, "tracked-address", false)
	if err != nil {
		return nil, err
	}

	log.Println("connected to levelDB")
	return db, nil
}

func connectToMongoDB() (*mongo.Client, error) {
	if os.Getenv("MONGODB_URI") == "" {
		return nil, errors.New("MONGODB_URI undefined")
	}

	if os.Getenv("MONGO_DATABASE") == "" {
		return nil, errors.New("undefined database")
	}

	if os.Getenv("MONGO_ROOT_USERNAME") == "" || os.Getenv("MONGO_ROOT_PASSWORD") == "" {
		return nil, errors.New("undefined credentials")
	}

	uri := os.Getenv("MONGODB_URI")

	clientOptions := options.Client().ApplyURI(uri)
	clientOptions.SetAuth(options.Credential{
		AuthSource: os.Getenv("MONGO_DATABASE"),
		Username:   os.Getenv("MONGO_ROOT_USERNAME"),
		Password:   os.Getenv("MONGO_ROOT_PASSWORD"),
	})

	mongoConn, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	log.Println("connected to mongoDB")
	return mongoConn, nil
}

func connectToGraphDB(mongoConn *mongo.Client) (*data.GraphDB, error) {
	if os.Getenv("MONGO_DATABASE") == "" {
		return nil, errors.New("undefined database")
	}

	dbName := os.Getenv("MONGO_DATABASE")
	log.Println("connected to graphDB")

	return &data.GraphDB{
		DbName: dbName,
		Db:     mongoConn.Database(dbName),
	}, nil
}

func connectToEthNode() (*ethclient.Client, error) {
	var NODE_URL string = GOERLI_WSS
	if os.Getenv("NODE_URL") != "" {
		NODE_URL = os.Getenv("NODE_URL")
	}

	conn, err := ethclient.Dial(NODE_URL)
	if err != nil {
		return nil, err
	}

	log.Println("connected to Ethereum node")
	return conn, nil
}
