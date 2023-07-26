package controllers

import (
	"github.com/dedpidgon/go-web-app/pkg/config"
	"github.com/dedpidgon/go-web-app/pkg/models"
	"github.com/dedpidgon/go-web-app/pkg/response"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type vote_controller struct{}

// /api/votes
func (vc *vote_controller) use(r *gin.RouterGroup) {
	r.POST("/:id", vc.vote)
}

func (vc *vote_controller) vote(c *gin.Context) {
	if !config.Client.Authenticated {
		response.Unauthorized(c, "you need to be authenticated to vote")
		return
	}
	var vote *models.Vote
	oid, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		response.BadReq(c, err)
		return
	}

	if err := c.ShouldBindJSON(&vote); err != nil {
		response.BadReq(c, err)
		return
	}

	vote.Parent = oid
	vote.Author = config.Client.ID

	if err := vote.DoTheThing(); err != nil {
		response.ServerErr(c, err)
		return
	}

	response.OK(c, "you are a voter!", vote)
}
