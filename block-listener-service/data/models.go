package data

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Vertex struct {
	Address   string    `json:"address" bson:"address"`
	Type      string    `json:"type" bson:"type"`
	TxEdges   []string  `json:"txEdges" bson:"txEdges"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

const (
	OutEdge = iota
	InEdge
)

type TxEdge struct {
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	Address     string    `json:"address" bson:"address"`
	Status      uint64    `json:"status" bson:"status"`
	TxHash      string    `json:"txHash" bson:"txHash"`
	BlockNumber string    `json:"blockNumber" bson:"blockNumber"`
	Value       string    `json:"value" bson:"value"`
	EventLog    []string  `json:"eventLog" bson:"eventLog"`
	Direct      int       `json:"direct" bson:"direct"`
}

type Event struct {
	Address string   `json:"address" bson:"address"`
	Topics  []string `json:"topics" bson:"topics"`
	Data    string   `json:"data" bson:"data"`
}

type GraphDB struct {
	DbName string
	Db     *mongo.Database
}

func (graphDB *GraphDB) GetCollection(collectionName string) *mongo.Collection {
	return graphDB.Db.Collection(collectionName)
}
