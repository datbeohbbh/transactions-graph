package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type IDAO interface {
	ExistedAddress(context.Context, string, string) (bool, error)
	GetAllVertex(context.Context) ([]*Vertex, error)
	GetVertexByAddress(context.Context, string) (*Vertex, error)
	GetVertexByAccountType(context.Context, string) ([]*Vertex, error)

	GetCollection(string) *mongo.Collection

	GetTxByAddress(context.Context, string) ([]*TxEdge, error)
	GetTxByObjectID(context.Context, string) (*TxEdge, error)
	GetTxByTxHash(context.Context, string) ([]*TxEdge, error)
	GetTxByBlockNumber(context.Context, string) ([]*TxEdge, error)
	GetTxByEdgeDirection(context.Context, int32) ([]*TxEdge, error)
}
