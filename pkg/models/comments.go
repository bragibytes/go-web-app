package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Author    primitive.ObjectID `json:"_author" bson:"_author"`
	Parent    primitive.ObjectID `json:"_parent" bson:"_parent"`
	Content   string             `json:"content" bson:"content" validate:"required" min:"3" max:"255"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
	Score     int32              `json:"score" bson:"score,omitempty"`
}

// crud
func (c *Comment) Save() error {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return err
	}
	res, err := comments_collection.InsertOne(ctx, c)
	if err != nil {
		return err
	}
	c.connect_to_parent()
	c.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
func (c *Comment) connect_to_parent() error {
	filter := bson.M{"_id": c.Parent}
	update := bson.M{"$push": bson.M{"comments": c.ID}}
	_, err := posts_collection.UpdateOne(ctx, filter, update)
	return err
}
func (c *Comment) disconnect_from_parent() error {
	filter := bson.M{"_id": c.Parent}
	update := bson.M{"$pull": bson.M{"comments": c.ID}}
	_, err := posts_collection.UpdateOne(ctx, filter, update)
	return err
}
func GetAllComments() ([]*Comment, error) {
	var comments []*Comment
	cursor, err := comments_collection.Find(ctx, bson.M{})
	if err != nil {
		return comments, err
	}
	err = cursor.All(ctx, &comments)
	if err != nil {
		return comments, err
	}
	for _, comment := range comments {
		calculate_score(comment)
	}
	return comments, nil
}
func GetOneComment(id primitive.ObjectID) (*Comment, error) {
	var comment *Comment
	err := comments_collection.FindOne(ctx, bson.M{"_id": id}).Decode(&comment)
	calculate_score(comment)
	return comment, err
}

func (c *Comment) Update(x *Post) error {
	err := comments_collection.FindOneAndUpdate(ctx, bson.M{"_id": c.ID}, bson.M{"$set": x}).Decode(&c)
	return err
}
func (c *Comment) Delete() error {
	_, err := posts_collection.DeleteOne(ctx, bson.M{"_id": c.ID})
	if err != nil {
		return err
	}
	err = c.disconnect_from_parent()
	return err
}

// voteable interface
func (c *Comment) add_to_score()              { c.Score += 1 }
func (c *Comment) subtract_from_score()       { c.Score -= 1 }
func (c *Comment) get_id() primitive.ObjectID { return c.ID }
