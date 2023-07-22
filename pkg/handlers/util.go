package handlers

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) error {

	if err := create_template_cache(); err != nil {
		return err
	}
	ustate = new_user_state()
	user_crud(r.Group("/users"))

	r.GET("/", ustate.soft_cookie_check, home_page)

	return nil
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
