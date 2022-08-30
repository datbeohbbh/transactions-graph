package graph

import (
	"log"

	"github.com/datbeohbbh/transactions-graph/graph-service/dao"

	"github.com/datbeohbbh/transactions-graph/graph-service/graphAlgo"
)

type GraphData struct {
	UnimplementedGraphDataServer

	dao dao.IDAO

	renderContext *graphAlgo.GraphRenderContext
}

func New(dao_ dao.IDAO) *GraphData {
	log.Println("connected to graph data handler")
	return &GraphData{
		dao:           dao_,
		renderContext: graphAlgo.New(),
	}
}
