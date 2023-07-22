package models

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.TODO()

var users *mongo.Collection
var posts *mongo.Collection

func Init(c *mongo.Client) {
	users = c.Database("theThing").Collection("users")
	posts = c.Database("theThing").Collection("posts")
}
