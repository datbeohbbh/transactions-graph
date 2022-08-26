package address

import (
	"context"
	"fmt"
)

func (addressHandler *AddressHandler) AddAddress(ctx context.Context, address *Address) (*Response, error) {
	addr, err := toEthereumAddress(address.Address)
	if err != nil {
		return nil, fmt.Errorf("error on convert to ethereum address: %v", err)
	}

	accountType, err := addressHandler.GetAccountType(ctx, addr)
	if err != nil {
		return nil, fmt.Errorf("error on get account type: %v", err)
	}

	status, err := addressHandler.dao.Add(ctx, addr, accountType)
	if err != nil {
		return nil, fmt.Errorf("error on add address %s: %v", addr, err)
	}

	Msg := fmt.Sprintf("address %s has been added and now it is being tracked", addr)
	if status == int(Response_TRACKED) {
		Msg = fmt.Sprintf("address %s has already been added", addr)
	}

	return &Response{
		Error:  false,
		Msg:    Msg,
		Status: Response_StatusCode(status),
	}, nil
}
