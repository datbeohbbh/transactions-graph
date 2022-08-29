package address

import "context"

func (addressClient *Client) AddAddress(ctx context.Context, address string) (bool, error) {
	resp, err := addressClient.client.AddAddress(ctx, &Address{Address: address})
	if err != nil {
		return false, err
	}
	return resp.Status == Response_ADDED, nil
}
