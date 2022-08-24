package graph

import (
	"block-listener/data"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type GraphData struct {
	UnimplementedGraphDataServer

	mongoClient *mongo.Client

	graphDB *data.GraphDB
}

func New(mongoConn *mongo.Client, graphDBInstance *data.GraphDB) *GraphData {
	log.Println("connected to graph data handler")
	return &GraphData{
		mongoClient: mongoConn,
		graphDB:     graphDBInstance,
	}
}

func (gd *GraphData) Close() {
	if err := gd.mongoClient.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
