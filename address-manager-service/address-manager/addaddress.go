package address

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func toEthereumAddress(address string) string {
	return common.HexToAddress(address).Hex()
}

func (addressHandler *AddressHandler) AddAddress(ctx context.Context, address *Address) (*Response, error) {
	addr := toEthereumAddress(address.Address)
	status, err := addressHandler.dao.Add(ctx, addr)
	if err != nil {
		return &Response{
			Error:  true,
			Msg:    "failed on add address",
			Status: Response_StatusCode(status),
		}, err
	}
	return &Response{
		Error:  false,
		Msg:    fmt.Sprintf("address %s has been added and now it is being tracked", addr),
		Status: Response_StatusCode(status),
	}, nil
}
