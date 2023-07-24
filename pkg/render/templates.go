package render

import (
	"fmt"
	"log"

	"github.com/dedpidgon/go-web-app/pkg/controllers"
	"github.com/dedpidgon/go-web-app/pkg/models"
	"github.com/dedpidgon/go-web-app/pkg/response"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Handle(r *gin.Engine) {
	if err := create_template_cache(); err != nil {
		log.Fatal(err.Error())
	}
	data = standard_stuff()

	r.GET("/", home_page)
	r.GET("/profile/:id", profile_page)
	r.GET("/profile", profile_page)
	r.GET("/board", board_page)
	r.GET("/about", about_page)
}

func home_page(c *gin.Context) {

	if err := render_template(c, "home"); err != nil {
		response.ServerErr(c, err)
		return
	}

}
func profile_page(c *gin.Context) {
	param := c.Param("id")
	var user *models.User
	var err error

	if param == "" {
		user, err = models.GetOneUser(controllers.UserController.UserID())
		if err != nil {
			response.NotFound(c, "user", err.Error())
			return
		}
		data.U = user
	} else {
		oid, err := primitive.ObjectIDFromHex(param)
		if err != nil {
			response.BadReq(c, err)
			return
		}
		user, err = models.GetOneUser(oid)
		if err != nil {
			response.NotFound(c, "user", err.Error())
			return
		}
		data.U = user
	}

	data.X = user.AsJsonString()
	if err := render_template(c, "profile"); err != nil {
		response.ServerErr(c, err)
		return
	}
}

func board_page(c *gin.Context) {
	fmt.Println(" ----- rendering board bage with .UserID of", controllers.UserController.UserID())
	if err := render_template(c, "board"); err != nil {
		response.ServerErr(c, err)
		return
	}
}

func post_page(c *gin.Context) {
	oid, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		response.BadReq(c, err)
		return
	}
	post, err := models.GetOnePost(oid)
	if err != nil {
		response.NotFound(c, "post", err.Error())
		return
	}
	data.P = post
	data.X = post.AsJsonString()
	if err := render_template(c, "post"); err != nil {
		response.ServerErr(c, err)
		return
	}
}

func about_page(c *gin.Context) {
	if err := render_template(c, "about"); err != nil {
		response.ServerErr(c, err)
		return
	}
}
