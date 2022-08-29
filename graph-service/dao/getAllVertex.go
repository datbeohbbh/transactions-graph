package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (dao *DAO) GetAllVertex(ctx context.Context) ([]*Vertex, error) {
	filter := bson.D{{}}
	cur, err := dao.GetCollection("vertex").Find(ctx, filter)
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
