package amqpClient

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (c *AmqpClient) Publish(ctx context.Context, exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error {
	err := c.channel.PublishWithContext(
		ctx,
		exchange,
		key,
		mandatory,
		immediate,
		msg,
	)
	return err
}
