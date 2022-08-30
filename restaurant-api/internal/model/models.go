package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      *string            `bson:"name" json:"name"`
	Phone     *string            `bson:"phone" json:"phone"`
	Password  *string            `bson:"password" json:"password,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type Shop struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      *string            `bson:"name" json:"name"`
	Logo      *string            `bson:"logo" json:"logo"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type Menu struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      *string            `bson:"name" json:"name"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type Food struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      *string            `bson:"name" json:"name"`
	Price     *float32           `bson:"price" json:"price"`
	ImageURL  *string            `bson:"image_url" json:"image_url,omitempty"`
	Stock     *int               `bson:"stock" json:"stock"`
	MenuID    *string            `bson:"meun_id" json:"meun_id"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type OrderMaster struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Amount     *float32           `bson:"amount" json:"amount"`
	TableID    *string            `bson:"table_id" json:"table_id"`
	Status     *string            `bson:"status" json:"status"`
	CustomerID *string            `bson:"customer_id" json:"customer_id"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}

type OrderDetail struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UnitPrice *float32           `bson:"unit_price" json:"unit_price"`
	OrderID   *string            `bson:"order_id" json:"order_id"`
	FoodID    *string            `bson:"food_id" json:"food_id"`
	FoodName  *string            `bson:"food_name" json:"food_name"`
	FoodImage *string            `bson:"food_image" json:"food_image"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type Table struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Code      *string            `bson:"code" json:"code"`
	MaxSeat   *int               `bson:"max_seat" json:"max_seat"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type Customer struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Phone     *string            `bson:"phone" json:"phone"`
	Username  *string            `bson:"username" json:"username"`
	Password  *string            `bson:"password" json:"password,omitempty"`
	Avatar    *string            `bson:"avatar" json:"avatar"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type Session struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID       string             `bson:"user_id" json:"user_id"`
	RefreshToken string             `bson:"refresh_token" json:"refresh_token"`
	UserAgent    string             `bson:"user_agent" josn:"user_agent"`
	ClientIP     string             `bson:"client_ip" json:"client_ip"`
	IsBlocked    bool               `bson:"is_blocked" json:"is_blocked"`
	ExpiredAt    time.Time          `bson:"expired_at" json:"expired_at"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
}
