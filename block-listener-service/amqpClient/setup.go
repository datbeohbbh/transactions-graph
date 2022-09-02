package amqpClient

import amqp "github.com/rabbitmq/amqp091-go"

func (c *AmqpClient) SetUp(exchange, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table) {
	c.channel.ExchangeDeclare(
		exchange,
		kind,
		durable,
		autoDelete,
		internal,
		noWait,
		args,
	)
}
