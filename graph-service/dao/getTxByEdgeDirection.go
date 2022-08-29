package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (dao *DAO) GetTxByEdgeDirection(ctx context.Context, direction int32) ([]*TxEdge, error) {
	filter := bson.D{{Key: "direct", Value: direction}}
	coll := dao.GetCollection("edge")

	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var result []*TxEdge
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		tx := TxEdge{}
		err = cursor.Decode(&tx)
		if err != nil {
			return nil, err
		}
		result = append(result, &tx)
	}
	return result, nil
}
