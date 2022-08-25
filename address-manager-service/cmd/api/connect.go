package main

import (
	"context"
	"errors"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
