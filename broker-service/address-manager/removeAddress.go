package address

import (
	"context"
)

func (addressClient *Client) RemoveAddress(ctx context.Context, address *Address) (*GinResponse, error) {
	resp, err := addressClient.client.RemoveAddress(ctx, address)
	if err != nil {
		return createResponse(true, Response_StatusCode_name[int32(Response_FAIL)], err.Error(), nil), err
	}
	return createResponse(false, Response_StatusCode_name[int32(resp.GetStatus())], resp.GetMsg(), nil), err
}
