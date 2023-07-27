package controllers

import (
	"github.com/dedpidgon/go-web-app/pkg/response"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var u *user_controller
var p *post_controller
var c *comment_controller

func Handle(r *gin.Engine) {
	u = &user_controller{}
	p = &post_controller{}
	c = &comment_controller{}

	u.use(r.Group("/api/users"))
	p.use(r.Group("/api/posts"))
	c.use(r.Group("/api/comments"))
}

func save_session(c *gin.Context, id primitive.ObjectID, name string) error {
	ses := sessions.Default(c)
	ses.Set("mongo_id", id)
	ses.Set("username", name)
	err := ses.Save()
	return err
}
func delete_session(c *gin.Context) error {
	ses := sessions.Default(c)
	ses.Clear()
	if err := ses.Save(); err != nil {
		response.ServerErr(c, err)
		return err
	}
	return nil
}
