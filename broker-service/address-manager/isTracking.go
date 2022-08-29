package address

import (
	"context"
)

func (addressClient *Client) IsTracking(ctx context.Context, address *Address) (*GinResponse, error) {
	resp, err := addressClient.client.IsTracking(ctx, address)
	if err != nil {
		return createResponse(true, Response_StatusCode_name[int32(Response_FAIL)], err.Error(), nil), err
	}
	return createResponse(false, Response_StatusCode_name[int32(resp.GetStatus())], resp.GetMsg(), nil), err
}
