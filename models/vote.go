package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Vote struct {
	Voter primitive.ObjectID `bson:"voter" json:"voter"`
	Value int8               `bson:"value" json:"value"` // -1 or 1
}
