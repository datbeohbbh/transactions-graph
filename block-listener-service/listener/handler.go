package listener

import (
	"block-listener/data"
	"context"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlockListener struct {
	UnimplementedBlockListenerServer

	ethClient *ethclient.Client

	mongoClient *mongo.Client

	graphDB *data.GraphDB
}

func New(ethConn *ethclient.Client, mongoConn *mongo.Client, graphDBInstance *data.GraphDB) *BlockListener {
	log.Println("connected to block listener handler")
	return &BlockListener{
		ethClient:   ethConn,
		mongoClient: mongoConn,
		graphDB:     graphDBInstance,
	}
}

func (bl *BlockListener) Close() {
	bl.ethClient.Close()
	if err := bl.mongoClient.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
