package dao

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
)

func (dao *DAO) ExistedAddress(ctx context.Context, coll, address string) (bool, error) {
	address = common.HexToAddress(address).Hex()
	filter := bson.D{{Key: "address", Value: address}}
	count, err := dao.GetCollection(coll).CountDocuments(ctx, filter)
	if err != nil {
		return false, fmt.Errorf("failed on check existance of address %s", address)
	}
	return (count > 0), nil
}
