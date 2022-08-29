package dao

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
)

func (dao *DAO) GetTxByAddress(ctx context.Context, address string) ([]*TxEdge, error) {
	address = common.HexToAddress(address).Hex()
	filter := bson.D{{Key: "address", Value: address}}
	coll := dao.GetCollection("edge")

	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	var result []*TxEdge
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
