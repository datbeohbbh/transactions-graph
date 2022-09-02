package api

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type AmqpApi interface {
	Connection
	GetChannel() error
	SetUp(exchange, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table)
	Publish(ctx context.Context, exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error
}
