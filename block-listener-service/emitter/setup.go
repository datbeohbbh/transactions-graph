package emitter

import amqp "github.com/rabbitmq/amqp091-go"

func (emitter *Emitter) SetUp(exchange, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table) {
	emitter.channel.ExchangeDeclare(
		exchange,
		kind,
		durable,
		autoDelete,
		internal,
		noWait,
		args,
	)
}
