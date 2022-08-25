package listener

import (
	"context"
	"encoding/json"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (bl *BlockListener) Start(ctx context.Context, exchange, key string) error {
	log.Println("Start listening new block")

	header := make(chan *types.Header)
	subs, err := bl.ethClient.SubscribeNewHead(ctx, header)
	if err != nil {
		return err
	}

	for {
		select {
		case err := <-subs.Err():
			return err
		case newHeader := <-header:
			go func() {
				log.Printf("received block %s, now publish it", newHeader.Number.String())
				err := bl.publishHeader(ctx, newHeader, exchange, key)
				if err != nil {
					log.Panicf("failed on publish header: %v", err)
				}
			}()
		}
	}
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
