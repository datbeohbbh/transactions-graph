package address

import (
	"context"
	"fmt"
)

func (addressHandler *AddressHandler) IsTracking(ctx context.Context, address *Address) (*Response, error) {
	addr, err := toEthereumAddress(address.Address)
	if err != nil {
		return nil, fmt.Errorf("error on convert to ethereum address: %v", err)
	}

	exist, err := addressHandler.dao.Exist(ctx, addr)
	if err != nil {
		return nil, fmt.Errorf("error on check if address %s exists: %v", addr, err)
	}

	status := Response_NOT_EXIST
	msg := fmt.Sprintf("address %s is not exists", addr)
	if exist {
		status = Response_TRACKED
		msg = fmt.Sprintf("address %s is being tracked", addr)
	}

	return &Response{
		Error:  false,
		Msg:    msg,
		Status: Response_StatusCode(status),
	}, nil
}
