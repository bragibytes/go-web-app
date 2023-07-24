package controllers

import (
	"github.com/dedpidgon/go-web-app/pkg/models"
	"github.com/dedpidgon/go-web-app/pkg/response"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type post_controller struct {
	errors []string
}

func new_post_controller() *post_controller {
	x := &post_controller{}
	return x
}

// /api/posts
func (pc *post_controller) use(r *gin.RouterGroup) {
	r.POST("/:id", pc.create)
	r.GET("/", pc.get_all)
	r.GET("/:id", pc.get_one)
	r.PUT("/:id", pc.update)
	r.DELETE("/:id", pc.delete)
}
func (pc *post_controller) create(c *gin.Context) {
	var post *models.Post
	oid, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		response.BadReq(c, err)
		return
	}
	if err := c.ShouldBindJSON(&post); err != nil {
		response.BadReq(c, err)
		return
	}
	post.Author = oid
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
	return
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
	if !UserController.IsAuthenticated() {
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
	if err = post.Update(postUpdate); err != nil {
		response.ServerErr(c, err)
		return
	}
	response.OK(c, "successfully updated post", post)
}
func (pc *post_controller) delete(c *gin.Context) {
	if !UserController.IsAuthenticated() {
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

// template data
func (pc *post_controller) GetAllPosts() []*models.Post {
	posts, err := models.GetAllPosts()
	if err != nil {
		pc.add_error(err)
		return nil
	}
	return posts
}
func (pc *post_controller) PostErrors() []string {
	return pc.errors
}
func (pc *post_controller) Comments() []*models.Comment {
	comments, err := models.GetAllComments()
	if err != nil {
		pc.add_error(err)
		return nil
	}
	return comments
}

// helpers
func (pc *post_controller) add_error(err error) {
	pc.errors = append(pc.errors, err.Error())
}
