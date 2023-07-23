package controllers

import "github.com/gin-gonic/gin"

type comment_controller struct{}

func new_comment_controller() *comment_controller {
	x := &comment_controller{}

	return x
}
func (cc *comment_controller) use(r *gin.RouterGroup) {

}

func (cc *comment_controller) create(c *gin.Context)   {}
func (cc *comment_controller) read(c *gin.Context)     {}
func (cc *comment_controller) read_all(c *gin.Context) {}
func (cc *comment_controller) update(c *gin.Context)   {}
func (cc *comment_controller) delete(c *gin.Context)   {}
