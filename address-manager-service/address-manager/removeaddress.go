package address

import (
	"context"
	"fmt"
)

func (addressHandler *AddressHandler) RemoveAddress(ctx context.Context, address *Address) (*Response, error) {
	addr := toEthereumAddress(address.Address)
	status, err := addressHandler.dao.Remove(ctx, addr)
	if err != nil {
		return &Response{
			Error:  true,
			Msg:    "failed on remove address",
			Status: Response_StatusCode(status),
		}, err
	}
	return &Response{
		Error:  false,
		Msg:    fmt.Sprintf("address %s has been removed", addr),
		Status: Response_StatusCode(status),
	}, nil
}
