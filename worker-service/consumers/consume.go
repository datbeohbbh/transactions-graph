package consumers

import amqp "github.com/rabbitmq/amqp091-go"

func (consumer *Consumer) Consume(consumerName string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan []byte, error) {
	queueName := consumer.queue.Name
	msg, err := consumer.channel.Consume(
		queueName,
		consumerName,
		autoAck,
		exclusive,
		noLocal,
		noWait,
		args,
	)
	if err != nil {
		return nil, err
	}

	data := make(chan []byte)
	go func() {
		defer close(data)
		for payload := range msg {
			data <- payload.Body
		}
	}()
	return data, nil
}
