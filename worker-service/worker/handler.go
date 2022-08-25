package worker

import (
	"context"
	"log"
	"worker/consumers"
	"worker/dao"

	addressManager "worker/address-manager"

	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/mongo"
)

type Worker struct {
	ethClient *ethclient.Client

	mongoClient *mongo.Client

	dao *dao.DAO

	addressClient *addressManager.Client

	consumer *consumers.Consumer
}

func New(ethConn *ethclient.Client, mongoConn *mongo.Client, dao *dao.DAO, client *addressManager.Client, consumer_ *consumers.Consumer) *Worker {
	log.Println("connected to block listener handler")
	return &Worker{
		ethClient:     ethConn,
		mongoClient:   mongoConn,
		dao:           dao,
		addressClient: client,
		consumer:      consumer_,
	}
}

func (worker *Worker) Close() {
	worker.ethClient.Close()
	if err := worker.mongoClient.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
