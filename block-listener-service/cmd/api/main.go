package main

import (
	"block-listener/emitter"
	"block-listener/listener"
	"context"
	"log"
	"os"
)

func failedOnError(msg string, err error) {
	if err != nil {
		log.Panicf("%s: %v", msg, err)
	}
}

func main() {
	ethClient, err := connectToEthereumNode(os.Getenv("NODE_URL"))
	failedOnError("failed on connect to Ethereum node", err)
	defer ethClient.Close()

	amqpConn, err := connectToRabbitMQ(
		os.Getenv("AMQP_USERNAME"),
		os.Getenv("AMQP_PASSWORD"),
		os.Getenv("AMQP_HOST"),
		os.Getenv("AMQP_PORT"))
	failedOnError("failed on connect to rabbitmq", err)
	defer amqpConn.Close()

	emitter, err := emitter.New(amqpConn)
	failedOnError("failed on create emitter instance", err)
	emitter.SetUp(
		os.Getenv("EXCHANGE_NAME"),
		os.Getenv("EXCHANGE_KIND"),
		false,
		true,
		false,
		false,
		nil)
	defer emitter.Close()

	blockListener := listener.New(ethClient, emitter)
	err = blockListener.Start(
		context.Background(),
		os.Getenv("EXCHANGE_NAME"),
		"block.header")
	failedOnError("failed on listen new block", err)
}
