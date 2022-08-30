package dbrepo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"lightsaid.com/restaurant-app/restaurant-api/internal/model"
)

func (m *mongoRepo) CreateAdmin(name, phone, password string) (*model.Admin, error) {
	tb := m.GetCollection("admin")

	ctx, cancel := context.WithTimeout(context.Background(), crudTimeout)
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

	admin.ID = id
	return &admin, nil
}

func (m *mongoRepo) GetAdminByID(id primitive.ObjectID) (*model.Admin, error) {
	tb := m.GetCollection("admin")

	ctx, cancel := context.WithTimeout(context.Background(), crudTimeout)
	defer cancel()

	var a model.Admin
	result := tb.FindOne(ctx, bson.M{"_id": id})

	err := result.Decode(&a)
	return &a, err
}

func (m *mongoRepo) GetAdminByPhone(phone string) (*model.Admin, error) {
	tb := m.GetCollection("admin")

	ctx, cancel := context.WithTimeout(context.Background(), crudTimeout)
	defer cancel()

	var a model.Admin
	result := tb.FindOne(ctx, bson.M{"phone": phone})

	err := result.Decode(&a)
	return &a, err
}
