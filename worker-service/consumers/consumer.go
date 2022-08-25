package consumers

import amqp "github.com/rabbitmq/amqp091-go"

type Consumer struct {
	channel *amqp.Channel
	queue   *amqp.Queue
}

func New(amqpConn *amqp.Connection) (*Consumer, error) {
	ch, err := amqpConn.Channel()
	if err != nil {
		return nil, err
	}
	return &Consumer{channel: ch}, nil
}

func (consumer *Consumer) Close() {
	consumer.channel.Close()
}
