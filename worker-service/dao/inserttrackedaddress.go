package dao

import (
	"context"
	"fmt"
	"time"
)

func (dao *DAO) InsertTrackedAddress(ctx context.Context, address, accountType string) (bool, error) {
	if has, err := dao.ExistedAddress(ctx, "tracking", address); err != nil {
		return false, fmt.Errorf("failed on check the existance of tracked address: %v", err)
	} else if has {
		return false, nil
	} else {
		trackedAddress := TrackedAddress{
			Address:   address,
			Type:      accountType,
			Timestamp: time.Now(),
		}
		_, err = dao.GetCollection("tracking").InsertOne(ctx, trackedAddress)
		if err != nil {
			return false, fmt.Errorf("failed on insert tracked address: %v", err)
		}
		return true, nil
	}
}
