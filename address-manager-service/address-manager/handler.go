package address

import (
	"address-manger/dao"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type AddressHandler struct {
	UnimplementedAddressManagerServer

	mongoClient *mongo.Client

	dao *dao.DAO
}

func failedOnError(msg string, err error) {
	if err != nil {
		panic(fmt.Errorf("error: %s - %v", msg, err))
	}
}

func New(_mongoClient *mongo.Client, _dao *dao.DAO) *AddressHandler {
	return &AddressHandler{
		mongoClient: _mongoClient,
		dao:         _dao,
	}
}

func (addressHandler *AddressHandler) Close() {
	if err := addressHandler.mongoClient.Disconnect(context.TODO()); err != nil {
		failedOnError("failed on disconnect mongoDB", err)
	}
}
