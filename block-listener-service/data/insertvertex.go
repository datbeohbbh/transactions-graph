package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (graphdb *GraphDB) InsertVertex(ctx context.Context, vertex *Vertex, txEdge *TxEdge, events []*Event) error {
	eventIds, err := graphdb.InsertEvent(ctx, events)
	if err != nil {
		return err
	}
	txEdge.EventLog = []string{}
	for _, eIds := range eventIds {
		txEdge.EventLog = append(txEdge.EventLog, eIds.(primitive.ObjectID).Hex())
	}

	edgeIds, err := graphdb.InsertTxEdge(ctx, txEdge)
	if err != nil {
		return err
	}

	vertexColl := graphdb.GetCollection("vertex")
	vertex.TxEdges = append(vertex.TxEdges, edgeIds.(primitive.ObjectID).Hex())
	vertex.CreatedAt = time.Now()
	vertex.UpdatedAt = time.Now()

	_, err = vertexColl.InsertOne(ctx, *vertex)
	if err != nil {
		return err
	}
	return nil
}
