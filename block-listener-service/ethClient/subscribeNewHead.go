package ethClient

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
)

func (c *EthNodeClient) SubscribeNewHead(ctx context.Context) <-chan *types.Header {
	// return c.client.SubscribeNewHead(ctx, header)
	header := make(chan *types.Header)
	subs, err := c.client.SubscribeNewHead(ctx, header)
	if err != nil {
		log.Panicf("failed on subscribe new header %+v", err)
	}
	headerData := make(chan *types.Header)
	go func() {
		defer close(headerData)
		for {
			select {
			case err = <-subs.Err():
				headerData <- nil
				return
			case newHeader := <-header:
				headerData <- newHeader
			}
		}
	}()
	return headerData
}
