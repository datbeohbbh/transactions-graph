package dao

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
)

func (dao *DAO) GetVertexByAddress(ctx context.Context, address string) (*Vertex, error) {
	address = common.HexToAddress(address).Hex()
	tracked, err := dao.ExistedAddress(ctx, "vertex", address)
	if err != nil {
		return nil, fmt.Errorf("failed on check if address %s exist: %v", address, err)
	}
	if !tracked {
		return nil, fmt.Errorf("address %s has not been tracked yet", address)
	}

	filter := bson.D{{"address", address}}
	v := Vertex{}
	err = dao.GetCollection("vertex").FindOne(ctx, filter).Decode(&v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
