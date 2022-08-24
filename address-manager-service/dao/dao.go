package dao

import "go.mongodb.org/mongo-driver/mongo"

type DAO struct {
	databaseName string

	database *mongo.Database
}

func New(_databaseName string, _database *mongo.Database) *DAO {
	return &DAO{
		databaseName: _databaseName,
		database:     _database,
	}
}

func (dao *DAO) GetCollection(coll string) *mongo.Collection {
	return dao.database.Collection(coll)
}
