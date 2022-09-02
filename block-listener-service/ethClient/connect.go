package ethClient

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

type EthNodeClient struct {
	client *ethclient.Client
}

var (
	NODE_URL = os.Getenv("NODE_URL")
)

func New() *EthNodeClient {
	return &EthNodeClient{}
}

func (c *EthNodeClient) GetClient() any {
	return c
}

func (c *EthNodeClient) Connect(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()
	if NODE_URL == "" {
		return errors.New("undefined node url")
	}

	conn, err := ethclient.DialContext(ctx, NODE_URL)
	if err != nil {
		return err
	}

	log.Println("connected to Ethereum node")
	c.client = conn
	return nil
}

func (c *EthNodeClient) Close() error {
	c.client.Close()
	return nil
}
