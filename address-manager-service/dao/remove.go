package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (dao *DAO) Remove(ctx context.Context, address string) (int, error) {
	exist, err := dao.Exist(ctx, address)
	if err != nil {
		return FAIL, err
	}
	if !exist {
		return NOT_EXIST, nil
	}
	filter := bson.D{{Key: "address", Value: address}}
	trackColl := dao.GetCollection("tracking")
	_, err = trackColl.DeleteOne(ctx, filter)
	if err != nil {
		return FAIL, err
	}
	return REMOVED, nil
}
