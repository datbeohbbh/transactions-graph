package worker

import (
	"github.com/datbeohbbh/transactions-graph/worker/dao"
	"github.com/ethereum/go-ethereum/core/types"
)

func processEvent(receipt *types.Receipt) []*dao.Event {
	events := []*dao.Event{}
	for _, log := range receipt.Logs {
		addr := log.Address.Hex()
		topics := []string{}
		for _, t := range log.Topics {
			topics = append(topics, t.Hex())
		}
		logData := string(log.Data)
		events = append(events, &dao.Event{
			Address: addr,
			Topics:  topics,
			Data:    logData,
		})
	}
	return events
}
