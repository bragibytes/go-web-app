package controllers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type posts struct{}

func (p posts) routes(r *gin.RouterGroup) {
	r.POST("/", p.create)
	r.GET("/{id}", p.read)
	r.GET("/", p.readAll)
	r.PUT("/{id}", tokens{}.check(), p.update)
	r.DELETE("/{id}", tokens{}.check(), p.delete)
}

// CRUD routes
func (p posts) create(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		response{false, "i couldnt mind that shit, get better data", err, nil, http.StatusBadRequest}.send(c)
		return
	}
	if err := post.Create(); err != nil {
		response{false, "i couldnt create that thing", err, nil, http.StatusInternalServerError}.send(c)
		return
	}
	response{true, "post created", nil, post, http.StatusCreated}.send(c)
}

func (p posts) read(c *gin.Context) {

}
func (p posts) readAll(c *gin.Context) {
	arr, err := models.Post{}.All()
	if err != nil {
		response{false, "i couldnt get them", err, nil, http.StatusInternalServerError}.send(c)
		return
	}
	response{true, "posts retrieved", nil, arr, http.StatusOK}.send(c)

}

func (p posts) update(c *gin.Context) {

}

func (p posts) delete(c *gin.Context) {

}
