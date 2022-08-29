package dao

import "go.mongodb.org/mongo-driver/mongo"

type DAO struct {
	dbName string
	db     *mongo.Database
}

func New(databaseName_ string, database_ *mongo.Database) *DAO {
	return &DAO{
		dbName: databaseName_,
		db:     database_,
	}
}

func (dao *DAO) GetCollection(collectionName string) *mongo.Collection {
	return dao.db.Collection(collectionName)
}
