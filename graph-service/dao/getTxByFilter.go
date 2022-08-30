package dao

import (
	"context"
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

func (dao *DAO) GetTxByFilter(ctx context.Context, conds map[string]string) ([]*TxEdge, error) {
	matchConds := []any{}
	for k, v := range conds {
		switch k {
		case "status":
			status, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				return nil, err
			}
			matchConds = append(matchConds, bson.D{{Key: k, Value: status}})
		case "direct":
			d, ok := DirectName[v]
			if !ok {
				return nil, fmt.Errorf("failed on parse direct: %s", v)
			}
			matchConds = append(matchConds, bson.D{{Key: k, Value: d}})
		default:
			matchConds = append(matchConds, bson.D{{Key: k, Value: v}})
		}
	}

	filters := bson.D{{
		Key:   "$and",
		Value: matchConds,
	}}
	coll := dao.GetCollection("edge")

	cursor, err := coll.Find(ctx, filters)
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
