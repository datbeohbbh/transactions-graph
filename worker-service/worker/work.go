package worker

import (
	context "context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
)

func (worker *Worker) Start(ctx context.Context) error {
	data, err := worker.consumer.Consume(
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	for {
		select {
		case <-time.After(5 * time.Minute):
			return fmt.Errorf("block process exceed time limit")
		case d := <-data:
			var header types.Header
			err = json.Unmarshal(d, &header)
			if err != nil {
				return err
			}
			startTime := time.Now()
			go func() {
				pctx := context.WithValue(ctx, "blockNumber", header.Number.String())
				err = worker.processBlock(pctx, &header)
				if err != nil {
					log.Panicf("failed on process block: %v", err)
				}
				done := time.Now()
				log.Printf("block #%s is completely processed in %v", header.Number, done.Sub(startTime))
			}()
		}
	}
}
