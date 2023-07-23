package controllers

import (
	"github.com/dedpidgon/go-web-app/pkg/models"
	"github.com/dedpidgon/go-web-app/pkg/response"
	"github.com/gin-gonic/gin"
)

type vote_controller struct{}

func new_vote_controller() *vote_controller {
	x := &vote_controller{}

	return x
}
func (vc *vote_controller) use(r *gin.RouterGroup) {
	r.POST("/", vc.vote)
}

func (vc *vote_controller) vote(c *gin.Context) {
	var vote *models.Vote
	if err := c.ShouldBindJSON(&vote); err != nil {
		response.BadReq(c, err)
		return
	}
	if err := vote.DoTheThing(); err != nil {
		response.ServerErr(c, err)
		return
	}
	response.OK(c, "you are a voter!", vote)
}
