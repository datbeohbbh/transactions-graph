package address

import (
	"address-manger/dao"
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/mongo"
)

type AddressHandler struct {
	UnimplementedAddressManagerServer

	mongoClient *mongo.Client

	dao *dao.DAO

	ethClient *ethclient.Client
}

func failedOnError(msg string, err error) {
	if err != nil {
		panic(fmt.Errorf("error: %s - %v", msg, err))
	}
}

func New(_mongoClient *mongo.Client, _dao *dao.DAO, ethClient_ *ethclient.Client) *AddressHandler {
	return &AddressHandler{
		mongoClient: _mongoClient,
		dao:         _dao,
		ethClient:   ethClient_,
	}
}

func (addressHandler *AddressHandler) Close() {
	if err := addressHandler.mongoClient.Disconnect(context.TODO()); err != nil {
		failedOnError("failed on disconnect mongoDB", err)
	}
	addressHandler.ethClient.Close()
}

func toEthereumAddress(address string) (string, error) {
	re, _ := regexp.Compile("0x[0-9a-zA-Z]{40}")
	if !re.MatchString(address) {
		return "", errors.New("provided address is not valid ethereum address")
	}
	return common.HexToAddress(address).Hex(), nil
}
