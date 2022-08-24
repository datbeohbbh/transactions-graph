package data

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (graphdb *GraphDB) GetVertexByAddress(ctx context.Context, address string, excludeFields any) ([]byte, error) {
	address = common.HexToAddress(address).Hex()
	tracked, err := graphdb.ExistedAddress(ctx, "vertex", address)
	if err != nil {
		return nil, fmt.Errorf("failed on check if address %s exist: %v", address, err)
	}
	if !tracked {
		return nil, fmt.Errorf("address %s has not been tracked yet", address)
	}

	filter := bson.D{{"address", address}}
	v := Vertex{}
	err = graphdb.GetCollection("vertex").FindOne(ctx, filter).Decode(&v)

	if err != nil {
		return nil, err
	}
	result, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (graphdb *GraphDB) GetAllVertex(ctx context.Context, excludeFields any) ([][]byte, error) {
	filter := bson.D{{}}
	opts := options.Find().SetProjection(excludeFields)
	cur, err := graphdb.GetCollection("vertex").Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	result := [][]byte{}
	for cur.Next(ctx) {
		v := Vertex{}
		err := cur.Decode(&v)
		if err != nil {
			return nil, err
		}
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		result = append(result, b)
	}
	return result, nil
}
