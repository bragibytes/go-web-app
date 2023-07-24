package render

import (
	"encoding/json"
	"github.com/dedpidgon/go-web-app/pkg/controllers"
	"github.com/dedpidgon/go-web-app/pkg/models"
	"github.com/dedpidgon/go-web-app/pkg/response"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"html/template"
	"log"
)

var path string
var cache map[string]*template.Template
var data *template_data

func Handle(r *gin.Engine) {
	path = "./templates/"
	cache = make(map[string]*template.Template)
	if err := create_template_cache(); err != nil {
		log.Fatal(err.Error())
	}
	r.GET("/", home_page)
	r.GET("/profile/:id", profile_page)
	r.GET("/profile", profile_page)
}

func home_page(c *gin.Context) {

	data = &template_data{
		controllers.UserController,
		controllers.PostController,
		nil,
		"",
	}
	if err := render_template(c.Writer, "home", data); err != nil {
		response.ServerErr(c, err)
		return
	}

}
func profile_page(c *gin.Context) {
	log.Println("in the profile_page handler")
	param := c.Param("id")
	if param == "" {
		user, err := models.GetOneUser(controllers.UserController.UserID())
		if err != nil {
			response.NotFound(c, "user", err.Error())
			return
		}
		data = &template_data{
			controllers.UserController,
			controllers.PostController,
			user,
			"",
		}
	} else {
		oid, err := primitive.ObjectIDFromHex(param)
		if err != nil {
			response.BadReq(c, err)
			return
		}
		user, err := models.GetOneUser(oid)
		if err != nil {
			response.NotFound(c, "user", err.Error())
			return
		}
		data = &template_data{
			controllers.UserController,
			controllers.PostController,
			user,
			"",
		}
	}
	bytes, err := json.Marshal(&data.U)
	if err != nil {
		response.Send(c, "info", err.Error(), bytes, 400)
		return
	}
	data.X = string(bytes)
	if err := render_template(c.Writer, "profile", data); err != nil {
		response.ServerErr(c, err)
		return
	}
}
