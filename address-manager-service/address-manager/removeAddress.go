package address

import (
	"context"
	"fmt"
)

func (addressHandler *AddressHandler) RemoveAddress(ctx context.Context, address *Address) (*Response, error) {
	addr, err := toEthereumAddress(address.GetAddress())
	if err != nil {
		return nil, fmt.Errorf("error on convert to ethereum address: %v", err)
	}

	status, err := addressHandler.dao.Remove(ctx, addr)
	if err != nil {
		return nil, fmt.Errorf("error on remove address: %v", err)
	}
	return &Response{
		Error:  false,
		Msg:    fmt.Sprintf("address %s has been removed", addr),
		Status: Response_StatusCode(status),
	}, nil
}
