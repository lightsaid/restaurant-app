package model

import "time"

type Shop struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name      *string   `bson:"name" josn:"name"`
	Logo      *string   `bson:"logo" json:"logo"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
