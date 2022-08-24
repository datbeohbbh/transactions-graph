package dao

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func (dao *DAO) ExistedAddress(ctx context.Context, coll, address string) (bool, error) {
	filter := bson.D{{"address", address}}
	count, err := dao.GetCollection(coll).CountDocuments(ctx, filter)
	if err != nil {
		return false, fmt.Errorf("failed on check existance of address %s", address)
	}
	return (count > 0), nil
}
