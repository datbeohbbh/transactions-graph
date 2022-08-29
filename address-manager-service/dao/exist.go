package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (dao *DAO) Exist(ctx context.Context, address string) (bool, error) {
	filter := bson.D{{Key: "address", Value: address}}
	trackColl := dao.GetCollection("tracking")

	counter, err := trackColl.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}
	return (counter > 0), nil
}
