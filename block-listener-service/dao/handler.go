package dao

import (
	"block-listener/data"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type DAO struct {
	UnimplementedDAOServer

	mongoClient *mongo.Client

	graphDB *data.GraphDB
}

func New(mongoConn *mongo.Client, graphDBInstance *data.GraphDB) *DAO {
	log.Println("connected to DAO handler")
	return &DAO{
		mongoClient: mongoConn,
		graphDB:     graphDBInstance,
	}
}

func (dao *DAO) Close() {
	if err := dao.mongoClient.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
