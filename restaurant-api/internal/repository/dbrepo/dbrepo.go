package dbrepo

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
	"lightsaid.com/restaurant-app/restaurant-api/internal/repository"
)

var (
	// ErrObjectIDToString primitive.ObjectID 转 string 出错了。
	ErrObjectIDToString = errors.New("InsertedID 转 String 失败。")
)

type mongoRepo struct {
	DB     *mongo.Client
	DBName string
}

func NewMongoRepo(client *mongo.Client, dbName string) repository.Repository {
	return &mongoRepo{
		DB:     client,
		DBName: dbName,
	}
}

func (m *mongoRepo) GetCollection(table string) *mongo.Collection {
	return m.DB.Database(m.DBName).Collection(table)
}
