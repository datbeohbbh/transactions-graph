package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (graphDB *GraphDB) UpdateVertex(ctx context.Context, vertex *Vertex, txEdge *TxEdge, events []*Event) error {
	eventIds, err := graphDB.InsertEvent(ctx, events)
	if err != nil {
		return err
	}
	txEdge.EventLog = []string{}
	for _, eIds := range eventIds {
		txEdge.EventLog = append(txEdge.EventLog, eIds.(primitive.ObjectID).Hex())
	}

	edgeIds, err := graphDB.InsertTxEdge(ctx, txEdge)
	if err != nil {
		return err
	}

	vertexColl := graphDB.GetCollection("vertex")

	filter := bson.D{{"address", vertex.Address}}
	pushEdgeIds := bson.D{{"$push", bson.D{{"txEdges", edgeIds.(primitive.ObjectID).Hex()}}}}
	update := bson.D{{"$set", bson.D{{"updatedAt", time.Now()}}}}

	if _, err = vertexColl.UpdateOne(ctx, filter, pushEdgeIds); err != nil {
		return err
	}
	if _, err = vertexColl.UpdateOne(ctx, filter, update); err != nil {
		return err
	}
	return nil
}
