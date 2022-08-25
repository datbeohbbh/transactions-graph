package address

import "context"

func (addressClient *Client) IsTracking(ctx context.Context, address string) (bool, error) {
	resp, err := addressClient.client.IsTracking(ctx, &Address{Address: address})
	if err != nil {
		return false, err
	}
	return resp.Status == Response_TRACKED, nil
}
