package address

import "context"

func (addressClient *Client) AddAddress(ctx context.Context, address string) (string, string, error) {
	resp, err := addressClient.client.AddAddress(ctx, &Address{Address: address})
	if err != nil {
		return Response_StatusCode_name[int32(Response_FAIL)], resp.Msg, err
	}
	status := resp.Status
	return Response_StatusCode_name[int32(status)], resp.Msg, nil
}
