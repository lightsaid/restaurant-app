package dbrepo

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
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

// GetCollection 获取集合
func (m *mongoRepo) GetCollection(table string) *mongo.Collection {
	return m.DB.Database(m.DBName).Collection(table)
}

// CreateIndex 创建索引
func (m *mongoRepo) CreateIndex(table string, field string, unique bool) bool {
	// 定义索引键
	mod := mongo.IndexModel{
		Keys:    bson.M{field: 1},
		Options: options.Index().SetUnique(unique),
	}

	// 超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// 操作的集合
	collection := m.GetCollection(table)

	_, err := collection.Indexes().CreateOne(ctx, mod)
	if err != nil {
		zap.S().Error(err)
		return false
	}

	return true
}
