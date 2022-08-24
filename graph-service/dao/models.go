package dao

import (
	"time"
)

type Vertex struct {
	Address   string    `json:"address" bson:"address"`
	Type      string    `json:"type,omitempty" bson:"type"`
	TxEdges   []string  `json:"txEdges,omitempty" bson:"txEdges"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updatedAt"`
}

const (
	OutEdge int = iota
	InEdge
)

type TxEdge struct {
	CreatedAt   time.Time `json:"createdAt,omitempty" bson:"createdAt"`
	Address     string    `json:"address" bson:"address"`
	Status      uint64    `json:"status,omitempty" bson:"status"`
	TxHash      string    `json:"txHash" bson:"txHash"`
	BlockNumber string    `json:"blockNumber,omitempty" bson:"blockNumber"`
	Value       string    `json:"value,omitempty" bson:"value"`
	EventLog    []string  `json:"eventLog,omitempty" bson:"eventLog"`
	Direct      int       `json:"direct,omitempty" bson:"direct"`
}

type Event struct {
	Address string   `json:"address" bson:"address"`
	Topics  []string `json:"topics,omitempty" bson:"topics"`
	Data    string   `json:"data,omitempty" bson:"data"`
}

type TrackedAddress struct {
	Address   string    `json:"address" bson:"address"`
	Type      string    `json:"type,omitempty" bson:"type"`
	Timestamp time.Time `json:"timestamp,omitempty" bson:"timestamp"`
}
