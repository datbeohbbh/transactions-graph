package listener

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (bl *BlockListener) Start(ctx context.Context, exchange, key string) error {
	log.Println("Start listening new block")

	for header := range bl.nodeClient.SubscribeNewHead(ctx) {
		if header == nil {
			return fmt.Errorf("failed on subscribe new header")
		}
		// closure problem. prevent a single header to be capture by all goroutines
		header := header
		go func() {
			log.Printf("received block %s, now publish it", header.Number.String())
			err := bl.publishHeader(ctx, header, exchange, key)
			if err != nil {
				log.Panicf("failed on publish header: %v", err)
			}
		}()
	}
	return nil
}

func (bl *BlockListener) publishHeader(ctx context.Context, header *types.Header, exchange, key string) error {
	payload, err := json.Marshal(*header)
	if err != nil {
		return err
	}
	err = bl.emitter.Publish(
		ctx,
		exchange,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        payload,
		})
	if err != nil {
		return err
	}
	return nil
}
