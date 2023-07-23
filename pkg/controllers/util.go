package controllers

import (
	"github.com/gin-gonic/gin"
)

var UserController *user_controller
var TemplateController *template_controller
var PostController *post_controller
var VoteController *vote_controller
var CommentController *comment_controller

func InitData() {
	TemplateController = new_template_controller()
	UserController = new_user_controller()
	PostController = new_post_controller()
	VoteController = new_vote_controller()
	CommentController = new_comment_controller()
}
func InitRoutes(r *gin.Engine) {
	VoteController.use(r.Group("/api/votes"))
	UserController.use(r.Group("/api/users"))
	PostController.use(r.Group("/api/posts"))
	CommentController.use(r.Group("/api/comments"))
	TemplateController.use(r)
}
