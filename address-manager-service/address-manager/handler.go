package address

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/datbeohbbh/transactions-graph/address-manger/dao"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type AddressHandler struct {
	UnimplementedAddressManagerServer

	dao dao.IDAO

	ethClient *ethclient.Client
}

func failedOnError(msg string, err error) {
	if err != nil {
		panic(fmt.Errorf("error: %s - %v", msg, err))
	}
}

func New(_dao dao.IDAO, ethClient_ *ethclient.Client) *AddressHandler {
	return &AddressHandler{
		dao:       _dao,
		ethClient: ethClient_,
	}
}

func toEthereumAddress(address string) (string, error) {
	re, _ := regexp.Compile("0x[0-9a-zA-Z]{40}")
	if !re.MatchString(address) {
		return "", errors.New("provided address is not valid ethereum address")
	}
	return common.HexToAddress(address).Hex(), nil
}
