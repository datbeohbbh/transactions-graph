package dao

import "go.mongodb.org/mongo-driver/mongo"

type DAO struct {
	DbName string
	Db     *mongo.Database
}

func New(databaseName_ string, database_ *mongo.Database) *DAO {
	return &DAO{
		DbName: databaseName_,
		Db:     database_,
	}
}

func (dao *DAO) GetCollection(collectionName string) *mongo.Collection {
	return dao.Db.Collection(collectionName)
}
