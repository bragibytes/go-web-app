package controllers

import (
	"fmt"

	"github.com/dedpidgon/go-web-app/pkg/config"
	"github.com/dedpidgon/go-web-app/pkg/models"
	"github.com/dedpidgon/go-web-app/pkg/response"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type comment_controller struct{}

func (cc *comment_controller) use(r *gin.RouterGroup) {
	r.POST("/", cc.create)
	r.GET("/", cc.get_all)
	r.GET("/{id}", cc.get_one)
	r.PUT("/{id}", cc.update)
	r.DELETE("/{id}", cc.delete)
	r.POST("/vote", cc.vote)
}
func (cc *comment_controller) create(c *gin.Context) {
	var comment *models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		response.BadReq(c, err)
		return
	}
	if err := comment.Save(); err != nil {
		response.ServerErr(c, err)
		return
	}
	response.Created(c, "comment", comment)
}
func (cc *comment_controller) get_one(c *gin.Context) {
	oid, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		response.BadReq(c, err)
		return
	}
	comment, err := models.GetOneComment(oid)
	if err != nil {
		response.NotFound(c, "comment", err.Error())
		return
	}
	response.OK(c, "successfully fetched the comment", comment)
}
func (cc *comment_controller) get_all(c *gin.Context) {
	comments, err := models.GetAllComments()
	if err != nil {
		response.NotFound(c, "comments", err.Error())
		return
	}
	response.OK(c, "successfully fetched comments", comments)

}
func (cc *comment_controller) update(c *gin.Context) {
	var comment_update *models.Comment
	oid, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		response.BadReq(c, err)
		return
	}
	comment := &models.Comment{ID: oid}
	if err := c.ShouldBindJSON(&comment_update); err != nil {
		response.BadReq(c, err)
	}

	response.OK(c, "successfully updated comment", comment)
}
func (cc *comment_controller) delete(c *gin.Context) {
	oid, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		response.BadReq(c, err)
		return
	}
	commentToDelete := &models.Comment{ID: oid}
	if err := commentToDelete.Delete(); err != nil {
		response.ServerErr(c, err)
		return
	}
	response.OK(c, "successfully deleted comment", nil)
}

func (cc *comment_controller) vote(c *gin.Context) {
	var vote *models.Vote
	if err := c.ShouldBindJSON(&vote); err != nil {
		response.BadReq(c, err)
		return
	}
	vote.Author = config.Client.ID
	comment := &models.Comment{ID: vote.Parent}
	if err := comment.Vote(vote); err != nil {
		response.ServerErr(c, err)
		return
	}

	fmt.Println("Vote success and it was a ", comment.OK)
	response.OK(c, comment.OK, nil)
}
