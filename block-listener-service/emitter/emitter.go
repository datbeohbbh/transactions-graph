package emitter

import amqp "github.com/rabbitmq/amqp091-go"

type Emitter struct {
	channel *amqp.Channel
}

func New(amqpConn_ *amqp.Connection) (*Emitter, error) {
	ch, err := amqpConn_.Channel()
	if err != nil {
		return nil, err
	}
	return &Emitter{channel: ch}, nil
}

func (emitter *Emitter) Close() {
	emitter.channel.Close()
}
