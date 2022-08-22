package repository

import "lightsaid.com/restaurant-app/restaurant-api/internal/model"

type Repository interface {
	CreateAdmin(name, phone, password string) (*model.Admin, error)

	CreateShop(name, logo string) (*model.Shop, error)
}
