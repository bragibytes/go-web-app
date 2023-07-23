package controllers

import (
	"github.com/gin-gonic/gin"
)

const Production bool = false

var UserController *user_controller
var TemplateController *template_controller

func InitData() {
	TemplateController = new_template_controller()
	UserController = new_user_controller()
}
func InitRoutes(r *gin.Engine) {
	UserController.use(r.Group("/users"))
	TemplateController.use(r)
}

type response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
}

func (r response) send(c *gin.Context) {
	c.JSON(r.Code, r)
}

func RunApp() {

}
