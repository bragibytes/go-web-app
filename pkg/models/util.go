package models

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.TODO()

var users_collection *mongo.Collection

func Init(c *mongo.Client) {
	users_collection = c.Database("theThing").Collection("users")
}
