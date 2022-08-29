package dao

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (dao *DAO) UpdateVertex(ctx context.Context, vertex *Vertex, txEdge *TxEdge, events []*Event) error {
	eventIds, err := dao.InsertEvent(ctx, events)
	if err != nil {
		return err
	}
	txEdge.EventLog = []string{}
	for _, eIds := range eventIds {
		txEdge.EventLog = append(txEdge.EventLog, eIds.(primitive.ObjectID).Hex())
	}

	edgeIds, err := dao.InsertTxEdge(ctx, txEdge)
	if err != nil {
		return err
	}

	vertexColl := dao.GetCollection("vertex")

	filter := bson.D{{Key: "address", Value: vertex.Address}}
	pushEdgeIds := bson.D{{Key: "$push", Value: bson.D{{Key: "txEdges", Value: edgeIds.(primitive.ObjectID).Hex()}}}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "updatedAt", Value: time.Now()}}}}

	if _, err = vertexColl.UpdateOne(ctx, filter, pushEdgeIds); err != nil {
		return err
	}
	if _, err = vertexColl.UpdateOne(ctx, filter, update); err != nil {
		return err
	}
	return nil
}
