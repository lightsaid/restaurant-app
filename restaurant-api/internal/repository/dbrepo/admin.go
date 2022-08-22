package dbrepo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"lightsaid.com/restaurant-app/restaurant-api/internal/model"
)

func (m *mongoRepo) CreateAdmin(name, phone, password string) (*model.Admin, error) {
	tb := m.GetCollection("admin")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	admin := model.Admin{
		Name:      &name,
		Phone:     &phone,
		Password:  &password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result, err := tb.InsertOne(ctx, admin)

	if err != nil {
		return &admin, err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return &admin, ErrObjectIDToString
	}

	admin.ID = id.Hex()
	return &admin, nil
}
