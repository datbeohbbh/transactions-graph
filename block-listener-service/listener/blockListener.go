package listener

import (
	"github.com/datbeohbbh/transactions-graph/block-listener/emitter"

	"github.com/ethereum/go-ethereum/ethclient"
)

type BlockListener struct {
	ethClient *ethclient.Client
	emitter   *emitter.Emitter
}

func New(ethClient_ *ethclient.Client, emitter_ *emitter.Emitter) *BlockListener {
	return &BlockListener{
		ethClient: ethClient_,
		emitter:   emitter_,
	}
}
