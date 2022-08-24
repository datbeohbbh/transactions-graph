package worker

import (
	"context"
	"log"
	"worker/dao"

	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/mongo"
)

type Worker struct {
	UnimplementedWorkerServer

	ethClient *ethclient.Client

	mongoClient *mongo.Client

	dao *dao.DAO
}

func New(ethConn *ethclient.Client, mongoConn *mongo.Client, dao *dao.DAO) *Worker {
	log.Println("connected to block listener handler")
	return &Worker{
		ethClient:   ethConn,
		mongoClient: mongoConn,
		dao:         dao,
	}
}

func (bl *Worker) Close() {
	bl.ethClient.Close()
	if err := bl.mongoClient.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
