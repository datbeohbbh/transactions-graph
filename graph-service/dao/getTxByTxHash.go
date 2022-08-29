package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (dao *DAO) GetTxByTxHash(ctx context.Context, txHash string) (*TxEdge, error) {
	filter := bson.D{{"txHash", txHash}}
	coll := dao.GetCollection("edge")

	result := TxEdge{}
	err := coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
