package dbrepo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"lightsaid.com/restaurant-app/restaurant-api/internal/model"
)

// CreateRefreshToken 创建
func (m *mongoRepo) CreateRefreshToken(t model.Session) (*model.Session, error) {
	tb := m.GetCollection("session")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.CreatedAt = time.Now()

	result, err := tb.InsertOne(ctx, t)

	if err != nil {
		return &t, err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return &t, ErrObjectIDToString
	}

	t.ID = id
	return &t, nil
}

// GetRefreshToken 获取
func (m *mongoRepo) GetRefreshToken(id primitive.ObjectID) (*model.Session, error) {
	tb := m.GetCollection("session")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result := tb.FindOne(ctx, bson.M{"_id": id})
	var s model.Session
	err := result.Decode(&s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
