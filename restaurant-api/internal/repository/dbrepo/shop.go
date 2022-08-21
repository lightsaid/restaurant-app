package dbrepo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"lightsaid.com/restaurant-app/restaurant-api/internal/model"
)

func (m *mongoRepo) CreateShop(name, logo string) (*model.Shop, error) {
	tb := m.GetCollection("shop")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	shop := model.Shop{
		Name:      &name,
		Logo:      &logo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result, err := tb.InsertOne(ctx, shop)

	if err != nil {
		return &shop, err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return &shop, ErrObjectIDToString
	}

	shop.ID = id.Hex()
	return &shop, nil
}
