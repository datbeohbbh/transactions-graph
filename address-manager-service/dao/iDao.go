package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type IDAO interface {
	Exist(context.Context, string) (bool, error)
	Add(context.Context, string, string) (int, error)
	Remove(context.Context, string) (int, error)
	GetCollection(string) *mongo.Collection
}
