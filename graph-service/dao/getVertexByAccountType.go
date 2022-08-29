package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (dao *DAO) GetVertexByAccountType(ctx context.Context, accountType string) ([]*Vertex, error) {
	filter := bson.D{{"type", accountType}}
	coll := dao.GetCollection("vertex")

	cur, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	var result []*Vertex
	for cur.Next(ctx) {
		v := Vertex{}
		err := cur.Decode(&v)
		if err != nil {
			return nil, err
		}
		result = append(result, &v)
	}
	return result, nil
}
