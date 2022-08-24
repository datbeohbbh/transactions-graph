package graph

import (
	"graph-service/dao"
	"log"
)

type GraphData struct {
	UnimplementedGraphDataServer

	dao *dao.DAO
}

func New(dao_ *dao.DAO) *GraphData {
	log.Println("connected to graph data handler")
	return &GraphData{dao: dao_}
}
