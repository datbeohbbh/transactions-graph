package address

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

const (
	USER_ACCOUNT     = "user"
	CONTRACT_ACCOUNT = "contract"
)

func (addressHandler *AddressHandler) GetAccountType(ctx context.Context, addr string) (string, error) {
	contractSize, err := addressHandler.ethClient.CodeAt(ctx, common.HexToAddress(addr), nil)
	if err != nil {
		return "", err
	}
	if len(contractSize) > 0 {
		return CONTRACT_ACCOUNT, nil
	}
	return USER_ACCOUNT, nil
}

func (addressHandler *AddressHandler) AccountType(ctx context.Context, address *Address) (*Type, error) {
	addr, err := toEthereumAddress(address.GetAddress())
	if err != nil {
		return nil, fmt.Errorf("error on convert to ethereum address: %v", err)
	}
	accountType, err := addressHandler.GetAccountType(ctx, addr)
	if err != nil {
		return nil, fmt.Errorf("error on get account type: %v", err)
	}
	return &Type{Type: accountType}, nil
}
