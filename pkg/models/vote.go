package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Vote struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Author    primitive.ObjectID `json:"_author" bson:"_author"`
	Parent    primitive.ObjectID `json:"_parent" bson:"-"`
	Value     int8               `json:"value" bson:"value"` // -1 or 1
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
