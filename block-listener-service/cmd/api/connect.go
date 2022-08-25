package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	amqp "github.com/rabbitmq/amqp091-go"
)

func connectToEthereumNode(nodeUrl string) (*ethclient.Client, error) {
	if nodeUrl == "" {
		return nil, errors.New("undefined node url")
	}

	conn, err := ethclient.Dial(nodeUrl)
	if err != nil {
		return nil, err
	}

	log.Println("connected to Ethereum node")
	return conn, nil
}

func connectToRabbitMQ(user, password, host, port string) (*amqp.Connection, error) {
	_, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	for {
		conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s", user, password, host, port))
		if err != nil {
			log.Println("failed on connect rabbitmq, reconnect...")
			time.Sleep(1 * time.Second)
		} else {
			log.Println("connected to rabbitmq")
			return conn, nil
		}
	}
}
