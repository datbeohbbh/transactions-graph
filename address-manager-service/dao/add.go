package dao

import (
	"context"
	"time"
)

func (dao *DAO) Add(ctx context.Context, address string) (int, error) {
	exist, err := dao.Exist(ctx, address)
	if err != nil {
		return FAIL, err
	}
	if exist {
		return TRACKED, nil
	}

	trackColl := dao.GetCollection("tracking")
	doc := TrackingAddress{
		Address:   address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = trackColl.InsertOne(ctx, doc)
	if err != nil {
		return FAIL, err
	}
	return ADDED, nil
}
