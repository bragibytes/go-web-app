package models

import (
	"context"
	"sort"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var DatabaseName string = "theThing"
var ctx = context.TODO()

var users_collection *mongo.Collection
var posts_collection *mongo.Collection
var comments_collection *mongo.Collection
var votes_collection *mongo.Collection

func Init(c *mongo.Client) {
	users_collection = c.Database(DatabaseName).Collection("users")
	posts_collection = c.Database(DatabaseName).Collection("posts")
	comments_collection = c.Database(DatabaseName).Collection("comments")
	votes_collection = c.Database(DatabaseName).Collection("votes")
}

type voteable interface {
	add_to_score()
	subtract_from_score()
	get_id() primitive.ObjectID
}

func sort_comments(comments []*Comment) {
	sort.Slice(comments, func(i, j int) bool {
		return comments[i].Score > comments[j].Score
	})
}
func sort_posts(posts []*Post) {
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Score > posts[j].Score
	})
}

func calculate_score(v voteable) error {
	filter := bson.M{"_parent": v.get_id()}
	var votes []*Vote
	cur, err := votes_collection.Find(ctx, filter)
	if err != nil {
		return err
	}
	if err := cur.All(ctx, &votes); err != nil {
		return err
	}
	for _, vote := range votes {
		if vote.IsUpvote {
			v.add_to_score()
		} else {
			v.subtract_from_score()
		}
	}
	return nil
}
