package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (dao *DAO) GetTxByObjectID(ctx context.Context, objectID string) (*TxEdge, error) {
	id, err := primitive.ObjectIDFromHex(objectID)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: id}}
	coll := dao.GetCollection("edge")

	result := TxEdge{}
	err = coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
