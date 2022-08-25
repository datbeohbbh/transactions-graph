package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func connectToMongoDB(mongoUri, mongoDatabase, mongoUsername, mongoPassword string) (*mongo.Client, error) {
	if mongoUri == "" {
		return nil, errors.New("undefined mongo uri")
	}

	if mongoDatabase == "" {
		return nil, errors.New("undefined database")
	}

	if mongoUsername == "" || mongoPassword == "" {
		return nil, errors.New("undefined credentials")
	}

	clientOptions := options.Client().ApplyURI(mongoUri)
	clientOptions.SetAuth(options.Credential{
		AuthSource: mongoDatabase,
		Username:   mongoUsername,
		Password:   mongoPassword,
	})

	mongoConn, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	log.Printf("user %s connected to mongo database %s", mongoUsername, mongoDatabase)
	return mongoConn, nil
}

func connectToDB(mongoConn *mongo.Client, dbName string) *mongo.Database {
	return mongoConn.Database(dbName)
}

func connectToEthereumNode(nodeUrl string) (*ethclient.Client, error) {
	if nodeUrl == "" {
		return nil, errors.New("undefined node url")
	}

	conn, err := ethclient.Dial(nodeUrl)
	if err != nil {
		return nil, err
	}

	log.Println("connected to Ethereum node")
	return conn, nil
}

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

func connectToRabbitMQ(user, password, host, port string) (*amqp.Connection, error) {
	_, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	for {
		conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s", user, password, host, port))
		if err != nil {
			log.Println("failed on connect rabbitmq, reconnect...")
			time.Sleep(1 * time.Second)
		} else {
			log.Println("connected to rabbitmq")
			return conn, nil
		}
	}
}
