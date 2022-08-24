package dao

import "time"

type TrackingAddress struct {
	Address   string    `json:"address" bson:"address"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

const (
	FAIL int = iota
	ADDED
	REMOVED
	TRACKED
	NOT_EXIST
)
