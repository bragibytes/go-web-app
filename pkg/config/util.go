package config

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Production bool
var Client *ClientData

func Init() {
	Production = false
	new_client()
}

func SetClientData() gin.HandlerFunc {
	return func(c *gin.Context) {
		ses := sessions.Default(c)
		id := ses.Get("mongo_id")
		name := ses.Get("username")
		if id == nil && name == nil {
			// not logged in
			Client.Authenticated = false
			Client.Name = "anonymous"
			Client.ID = primitive.NilObjectID
		} else {
			// logged in
			Client.Authenticated = true
			Client.Name = name.(string)
			Client.ID = id.(primitive.ObjectID)
		}

	}
}
