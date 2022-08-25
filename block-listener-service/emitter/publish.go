package emitter

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (emitter *Emitter) Publish(ctx context.Context, exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error {
	err := emitter.channel.PublishWithContext(
		ctx,
		exchange,
		key,
		mandatory,
		immediate,
		msg,
	)
	return err
}
