package api

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
)

type EthApi interface {
	Connection
	SubscribeNewHead(ctx context.Context) <-chan *types.Header
}
