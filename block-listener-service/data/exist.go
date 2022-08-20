package data

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func (graphDB *GraphDB) ExistedVertex(ctx context.Context, address string) (bool, error) {
	filter := bson.D{{"address", address}}
	count, err := graphDB.GetCollection("vertex").CountDocuments(ctx, filter)
	if err != nil {
		return false, fmt.Errorf("failed on check existance of address %s", address)
	}
	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
