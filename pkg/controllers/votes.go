package controllers

import (
	"github.com/dedpidgon/go-web-app/pkg/models"
	"github.com/dedpidgon/go-web-app/pkg/response"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type vote_controller struct{}

func new_vote_controller() *vote_controller {
	x := &vote_controller{}

	return x
}

// /api/votes
func (vc *vote_controller) use(r *gin.RouterGroup) {
	r.POST("/:id", vc.vote)
}

func (vc *vote_controller) vote(c *gin.Context) {
	var vote *models.Vote
	oid, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		response.ServerErr(c, err)
		return
	}
	if err := c.ShouldBindJSON(&vote); err != nil {
		response.BadReq(c, err)
		return
	}
	vote.Parent = oid
	vote.Author = UserController.UserID()
	if err := vote.DoTheThing(); err != nil {
		response.ServerErr(c, err)
		return
	}
	response.OK(c, "you are a voter!", vote)
}
