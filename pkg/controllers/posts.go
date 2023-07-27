package controllers

import (
	"github.com/dedpidgon/go-web-app/pkg/config"
	"github.com/dedpidgon/go-web-app/pkg/models"
	"github.com/dedpidgon/go-web-app/pkg/response"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type post_controller struct {
}

// /api/posts
func (pc *post_controller) use(r *gin.RouterGroup) {
	r.POST("/", pc.create)
	r.GET("/", pc.get_all)
	r.GET("/:id", pc.get_one)
	r.PUT("/:id", pc.update)
	r.DELETE("/:id", pc.delete)
	r.POST("/vote", pc.vote)
}
func (pc *post_controller) create(c *gin.Context) {
	var post *models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		response.BadReq(c, err)
		return
	}
	if validation_errors := post.Valid(); len(validation_errors) > 0 {
		response.ValidationErrors(c, validation_errors)
		return
	}
	if err := post.Save(); err != nil {
		response.ServerErr(c, err)
		return
	}
	response.Created(c, "post", post)
}
func (pc *post_controller) get_one(c *gin.Context) {
	oid, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		response.BadReq(c, err)
		return
	}

	post, err := models.GetOnePost(oid)
	if err != nil {
		response.ServerErr(c, err)
		return
	}
	response.OK(c, "got the post", post)

}
func (pc *post_controller) get_all(c *gin.Context) {

	posts, err := models.GetAllPosts()
	if err != nil {
		response.NotFound(c, "post", err.Error())
		return
	}
	response.OK(c, "got all the posts", posts)

}
func (pc *post_controller) update(c *gin.Context) {
	if !config.Client.Authenticated {
		response.Unauthorized(c, "i cant let you do that")
		return
	}
	var post *models.Post
	var postUpdate *models.Post
	oid, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		response.BadReq(c, err)
		return
	}
	post, err = models.GetOnePost(oid)
	if err != nil {
		response.NotFound(c, "post", err.Error())
		return
	}
	if err := c.ShouldBindJSON(&postUpdate); err != nil {
		response.BadReq(c, err)
		return
	}

	postUpdate.UpdateFriend(post)
	if validation_errors := post.Valid(); len(validation_errors) > 0 {
		response.ValidationErrors(c, validation_errors)
		return
	}
	if err = post.Update(); err != nil {
		response.ServerErr(c, err)
		return
	}
	response.OK(c, "successfully updated post", post)
}
func (pc *post_controller) delete(c *gin.Context) {
	if !config.Client.Authenticated {
		response.Unauthorized(c, "i cant let you do that")
		return
	}
	oid, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		response.BadReq(c, err)
		return
	}
	postToDelete := &models.Post{ID: oid}
	if err := postToDelete.Delete(); err != nil {
		response.ServerErr(c, err)
		return
	}
	response.OK(c, "successfully deleted post", nil)
}

func (p *post_controller) vote(c *gin.Context) {
	var vote *models.Vote
	if err := c.ShouldBindJSON(&vote); err != nil {
		response.BadReq(c, err)
		return
	}
	vote.Author = config.Client.ID
	post := &models.Post{ID: vote.Parent}
	if err := post.Vote(vote); err != nil {
		response.ServerErr(c, err)
		return
	}

	response.OK(c, "thanks for voting", nil)
}
