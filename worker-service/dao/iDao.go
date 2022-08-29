package dao

import (
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type IDAO interface {
	ExistedAddress(context.Context, string, string) (bool, error)
	InsertEvent(context.Context, []*Event) ([]any, error)
	InsertTxEdge(context.Context, *TxEdge) (any, error)
	InsertVertex(context.Context, *Vertex, *TxEdge, []*Event) error
	UpdateVertex(context.Context, *Vertex, *TxEdge, []*Event) error
	UpdateDB(context.Context, *Vertex, *TxEdge, []*Event) error
	GetCollection(string) *mongo.Collection
}
