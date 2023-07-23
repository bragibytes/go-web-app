package controllers

import (
	"github.com/dedpidgon/go-web-app/pkg/models"
	"github.com/dedpidgon/go-web-app/pkg/response"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type comment_controller struct{}

func new_comment_controller() *comment_controller {
	x := &comment_controller{}

	return x
}

func (cc *comment_controller) use(r *gin.RouterGroup) {
	r.POST("/", cc.create)
	r.GET("/", cc.get_all)
	r.GET("/{id}", cc.get_one)
	r.PUT("/{id}", cc.update)
	r.DELETE("/{id}", cc.delete)
}
func (cc *comment_controller) create(c *gin.Context) {
	var comment *models.Comment
	if err := c.ShouldBindJSON(&cc); err != nil {
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
	return
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
