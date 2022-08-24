package address

import (
	"context"
	"fmt"
)

func (addressHandler *AddressHandler) IsTracking(ctx context.Context, address *Address) (*Response, error) {
	addr := toEthereumAddress(address.Address)
	exist, err := addressHandler.dao.Exist(ctx, addr)
	if err != nil {
		return &Response{
			Error:  true,
			Msg:    "failed on checking if address exists",
			Status: Response_FAIL,
		}, err
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
