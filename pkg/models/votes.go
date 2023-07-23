package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Vote struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Author    primitive.ObjectID `json:"_author" bson:"_author"`
	Parent    primitive.ObjectID `json:"_parent" bson:"_parent"`
	IsUpvote  bool               `json:"is_upvote" bson:"is_upvote"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func (v *Vote) DoTheThing() error {

	var xvote *Vote
	// check if the vote already exists
	filter := bson.M{"_author": v.Author, "_parent": v.Parent}
	if err := votes_collection.FindOne(ctx, filter).Decode(&xvote); err != nil {
		// does not exist
		v.CreatedAt = time.Now()
		v.UpdatedAt = time.Now()
		res, err := votes_collection.InsertOne(ctx, v)
		if err != nil {
			return err
		}
		v.ID = res.InsertedID.(primitive.ObjectID)
	} else {
		if v.IsUpvote == xvote.IsUpvote {
			// clicked the same vote button, delete the vote
			if err := xvote.delete(); err != nil {
				return err
			}
		} else {
			update := bson.M{"$set": bson.M{"is_upvote": !v.IsUpvote, "updated_at": time.Now()}}
			filter := bson.M{"_id": v.ID}
			_, err := votes_collection.UpdateOne(ctx, filter, update)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

func (v *Vote) delete() error {
	filter := bson.M{"_id": v.ID}
	_, err := votes_collection.DeleteOne(ctx, filter)
	return err
}
