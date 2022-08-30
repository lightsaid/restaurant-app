package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"lightsaid.com/restaurant-app/restaurant-api/internal/model"
)

type Repository interface {
	CreateAdmin(name, phone, password string) (*model.Admin, error)
	GetAdminByID(id primitive.ObjectID) (*model.Admin, error)
	GetAdminByPhone(phone string) (*model.Admin, error)

	CreateShop(name, logo string) (*model.Shop, error)
	GetShops() ([]*model.Shop, error)

	GetRefreshToken(id primitive.ObjectID) (*model.Session, error)
	CreateRefreshToken(t model.Session) (*model.Session, error)
}
