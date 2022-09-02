package main

import (
	"context"
	"log"
	"os"

	"github.com/datbeohbbh/transactions-graph/block-listener/amqpClient"
	"github.com/datbeohbbh/transactions-graph/block-listener/emitter"
	"github.com/datbeohbbh/transactions-graph/block-listener/ethClient"
	api "github.com/datbeohbbh/transactions-graph/block-listener/interfaces"
	"github.com/datbeohbbh/transactions-graph/block-listener/listener"
)

func failedOnError(msg string, err error) {
	if err != nil {
		log.Panicf("%s: %v", msg, err)
	}
}

func main() {
	var blockChainNetworkClient api.EthApi = ethClient.New()
	err := blockChainNetworkClient.Connect(context.Background())
	failedOnError("failed on connect to node client", err)
	defer blockChainNetworkClient.Close()

	var messageBrokerClient api.AmqpApi = amqpClient.New()
	err = messageBrokerClient.Connect(context.Background())
	failedOnError("failed on connect to rabbitmq", err)
	defer messageBrokerClient.Close()

	emitter, err := emitter.New(messageBrokerClient)
	failedOnError("failed on create emitter instance", err)
	emitter.SetUp(
		os.Getenv("EXCHANGE_NAME"),
		os.Getenv("EXCHANGE_KIND"),
		false,
		true,
		false,
		false,
		nil)

	blockListener := listener.New(blockChainNetworkClient, emitter)
	err = blockListener.Start(
		context.Background(),
		os.Getenv("EXCHANGE_NAME"),
		"block.header")
	failedOnError("failed on listen new block", err)
}
