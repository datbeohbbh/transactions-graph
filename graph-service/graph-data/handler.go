package graph

import (
	"log"

	"github.com/datbeohbbh/transactions-graph/graph-service/dao"
)

type GraphData struct {
	UnimplementedGraphDataServer

	dao dao.IDAO
}

func New(dao_ dao.IDAO) *GraphData {
	log.Println("connected to graph data handler")
	return &GraphData{dao: dao_}
}
