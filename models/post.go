package models

import (
	"context"
	"time"

	_ "github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" validate:"required"`
	Title     string             `bson:"title,omitempty" validate:"required, min=3, max=50"`
	Content   string             `bson:"content,omitempty" validate:"required, min=10, max=1000"`
	Author    primitive.ObjectID `bson:"_author,omitempty" validate:"required"`
	CreatedAt time.Time          `bson:"created_at,omitempty" validate:"required"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" validate:"required"`

	Downvoted []primitive.ObjectID `bson:"downvoted,omitempty"`
	Upvoted   []primitive.ObjectID `bson:"upvoted,omitempty"`
	Votes     []Vote               `bson:"votes,omitempty"`
	Comments  []Comment            `bson:"comments,omitempty"`
}

// Create
func (p *Post) Create() error {

	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	_, err := posts.InsertOne(context.Background(), p)
	return err
}

// Read
func (p Post) All() ([]*Post, error) {
	var x []*Post
	cur, err := posts.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var n *Post
		if err := cur.Decode(&n); err != nil {
			return nil, err
		}
		x = append(x, n)
	}

	return x, nil

}
func (p Post) Read() (*Post, error) {
	return nil, nil
}

// Update
func (p *Post) Update(update *Post) error {
	_, err := posts.UpdateOne(ctx, bson.M{"_id": p.ID}, bson.M{"$set": update})
	return err
}

// Delete
func (p *Post) Delete() error {
	_, err := posts.DeleteOne(ctx, bson.M{"_id": p.ID})
	return err
}
