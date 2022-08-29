package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type IDAO interface {
	ExistedAddress(context.Context, string, string) (bool, error)
	GetAllVertex(context.Context) ([]*Vertex, error)
	GetVertexByAddress(context.Context, string) (*Vertex, error)
	GetCollection(string) *mongo.Collection
}
