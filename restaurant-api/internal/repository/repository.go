package repository

import "lightsaid.com/restaurant-app/restaurant-api/internal/model"

type Repository interface {
	CreateShop(name, logo string) (*model.Shop, error)
}
