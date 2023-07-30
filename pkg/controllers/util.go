package controllers

import (
	"github.com/dedpidgon/go-web-app/pkg/response"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var user_control *user_controller
var post_control *post_controller
var comment_control *comment_controller
var view_control *view_controller
var email_control *email_controller
var token_control *token_controller

func Handle(r *gin.Engine) {

	user_control = &user_controller{}
	post_control = &post_controller{}
	comment_control = &comment_controller{}
	view_control = &view_controller{}
	email_control = new_email_controller()
	token_control = new_token_controller()

	user_control.use(r.Group("/api/users"))
	post_control.use(r.Group("/api/posts"))
	comment_control.use(r.Group("/api/comments"))
	view_control.use(r)
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
