package emitter

import (
	api "github.com/datbeohbbh/transactions-graph/block-listener/interfaces"
)

type Emitter struct {
	amqpClient api.AmqpApi
}

func New(amqpApi api.AmqpApi) (*Emitter, error) {
	err := amqpApi.GetChannel()
	if err != nil {
		return nil, err
	}
	return &Emitter{amqpClient: amqpApi}, nil
}
