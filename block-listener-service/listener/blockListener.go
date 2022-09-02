package listener

import (
	"github.com/datbeohbbh/transactions-graph/block-listener/emitter"
	api "github.com/datbeohbbh/transactions-graph/block-listener/interfaces"
)

type BlockListener struct {
	nodeClient api.EthApi
	emitter    *emitter.Emitter
}

func New(ethClient api.EthApi, emitter_ *emitter.Emitter) *BlockListener {
	return &BlockListener{
		nodeClient: ethClient,
		emitter:    emitter_,
	}
}
