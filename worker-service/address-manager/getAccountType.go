package address

import "context"

func (addressClient *Client) AccountType(ctx context.Context, address string) (string, error) {
	accType, err := addressClient.client.AccountType(ctx, &Address{Address: address})
	if err != nil {
		return "", err
	}
	return accType.Type, nil
}
