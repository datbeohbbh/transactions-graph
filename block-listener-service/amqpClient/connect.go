package amqpClient

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type AmqpClient struct {
	client  *amqp.Connection
	channel *amqp.Channel
}

func New() *AmqpClient {
	return &AmqpClient{}
}

func (c *AmqpClient) Connect(ctx context.Context) error {
	var (
		user     = os.Getenv("AMQP_USERNAME")
		password = os.Getenv("AMQP_PASSWORD")
		host     = os.Getenv("AMQP_HOST")
		port     = os.Getenv("AMQP_PORT")
	)
	_, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()

	for {
		conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s", user, password, host, port))
		if err != nil {
			log.Println("failed on connect rabbitmq, reconnect...")
			time.Sleep(1 * time.Second)
		} else {
			log.Println("connected to rabbitmq")
			c.client = conn
			return nil
		}
	}
}

func (c *AmqpClient) Close() error {
	err := c.client.Close()
	if err != nil {
		return err
	}

	err = c.channel.Close()
	if err != nil {
		return err
	}
	return nil
}

func (c *AmqpClient) GetClient() any {
	return c
}
