package data

import (
	"context"
	"fmt"
	"time"
)

func (graphdb *GraphDB) InsertTrackedAddress(ctx context.Context, address, accountType string) (bool, error) {
	if has, err := graphdb.ExistedAddress(ctx, "tracking", address); err != nil {
		return false, fmt.Errorf("failed on check the existance of tracked address: %v", err)
	} else if has {
		return false, nil
	} else {
		trackedAddress := TrackedAddress{
			Address:   address,
			Type:      accountType,
			Timestamp: time.Now(),
		}
		_, err = graphdb.GetCollection("tracking").InsertOne(ctx, trackedAddress)
		if err != nil {
			return false, fmt.Errorf("failed on insert tracked address: %v", err)
		}
		return true, nil
	}
}
