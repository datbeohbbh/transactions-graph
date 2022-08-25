package consumers

import (
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

// this implementation is not good
// it is not good for large scale
// should learn more on design pattern.
func (consumer *Consumer) SetUp() error {
	consumer.SetUpExchange(
		os.Getenv("EXCHANGE_NAME"),
		os.Getenv("EXCHANGE_KIND"),
		false,
		true,
		false,
		false,
		nil)

	err := consumer.SetUpQueue(
		"block-header-queue",
		false,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = consumer.BindQueueWithExchange(
		"block.header",
		os.Getenv("EXCHANGE_NAME"),
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = consumer.channel.Qos(
		1,
		0,
		false,
	)
	if err != nil {
		return err
	}

	return nil
}

func (consumer *Consumer) SetUpExchange(exchange, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table) {
	consumer.channel.ExchangeDeclare(
		exchange,
		kind,
		durable,
		autoDelete,
		internal,
		noWait,
		args,
	)
}

func (consumer *Consumer) SetUpQueue(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) error {
	que, err := consumer.channel.QueueDeclare(
		name,
		durable,
		autoDelete,
		exclusive,
		noWait,
		args,
	)
	if err != nil {
		return err
	}
	consumer.queue = &que
	return nil
}

func (consumer *Consumer) BindQueueWithExchange(key, exchange string, noWait bool, args amqp.Table) error {
	queueName := consumer.queue.Name
	err := consumer.channel.QueueBind(
		queueName,
		key,
		exchange,
		noWait,
		args,
	)
	return err
}
