package controllers

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.RouterGroup) {
	UserController{}.routes(r.Group("/users"))
	posts{}.routes(r.Group("/posts"))
}

type response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   error       `json:"error"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
}

func (r response) send(c *gin.Context) {
	c.JSON(r.Code, r)
}
