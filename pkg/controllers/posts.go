package controllers

import "github.com/gin-gonic/gin"

type post_controller struct {
}

func new_post_controller() *post_controller {
	x := &post_controller{}
	return x
}

func (c *post_controller) use(r *gin.RouterGroup) {

}
func (pc *post_controller) create(c *gin.Context)   {}
func (pc *post_controller) read(c *gin.Context)     {}
func (pc *post_controller) read_all(c *gin.Context) {}
func (pc *post_controller) update(c *gin.Context)   {}
func (pc *post_controller) delete(c *gin.Context)   {}
