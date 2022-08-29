package worker

import (
	"log"

	"github.com/datbeohbbh/transactions-graph/worker/consumers"
	"github.com/datbeohbbh/transactions-graph/worker/dao"

	addressManager "github.com/datbeohbbh/transactions-graph/worker/address-manager"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Worker struct {
	ethClient *ethclient.Client

	dao dao.IDAO

	addressClient *addressManager.Client

	consumer *consumers.Consumer
}

func New(ethConn *ethclient.Client, dao *dao.DAO, client *addressManager.Client, consumer_ *consumers.Consumer) *Worker {
	log.Println("connected to block listener handler")
	return &Worker{
		ethClient:     ethConn,
		dao:           dao,
		addressClient: client,
		consumer:      consumer_,
	}
}
