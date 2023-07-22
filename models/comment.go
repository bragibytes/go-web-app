package models

import (
	_ "github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" validate:"required"`
	Author    primitive.ObjectID `bson:"author,omitempty" json:"author,omitempty" validate:"required"`
	Content   string             `bson:"content,omitempty" json:"content,omitempty" validate:"required, min=3, max=120"`
	CreatedAt string             `bson:"created_at,omitempty" json:"created_at,omitempty" validate:"required"`
	UpdatedAt string             `bson:"updated_at,omitempty" json:"updated_at,omitempty" validate:"required"`
}

// Create
func (c *Comment) Create() error {
	return nil
}

// Read
func (c *Comment) All() ([]Comment, error) {
	return nil, nil
}
func (c *Comment) Read() (*Comment, error) {
	return nil, nil
}

// Update
func (c *Comment) Update() error {
	return nil
}

// Delete
func (c *Comment) Delete() error {
	return nil
}
