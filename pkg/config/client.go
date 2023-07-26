package config

import "go.mongodb.org/mongo-driver/bson/primitive"

type ClientData struct {
	Authenticated bool
	Name          string
	ID            primitive.ObjectID
	IP            string
}

func new_client() {
	x := &ClientData{
		Authenticated: false,
		Name:          "anonymous",
		ID:            primitive.NilObjectID,
		IP:            "",
	}
	Client = x
}
