package address

import (
	"context"
	"fmt"
)

func (addressHandler *AddressHandler) AddAddress(ctx context.Context, address *Address) (*Response, error) {
	addr := toEthereumAddress(address.Address)
	accountType, err := addressHandler.GetAccountType(ctx, addr)

	if err != nil {
		return &Response{
			Error:  true,
			Msg:    "failed on get account type",
			Status: Response_FAIL,
		}, err
	}

	status, err := addressHandler.dao.Add(ctx, addr, accountType)
	if err != nil {
		return &Response{
			Error:  true,
			Msg:    "failed on add address",
			Status: Response_StatusCode(status),
		}, err
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
