package forum

import (
	v1 "forum-server/api/v1"

	"github.com/gin-gonic/gin"
)

type VoteRouter struct{}

func (v *VoteRouter) InitVoteRouter(router *gin.RouterGroup) {
	voteRouter := router.Group("vote")
	frmVoteApi := v1.ApiGroupApp.ForumApiGroup.VoteApi
	{
		voteRouter.POST("createVote", frmVoteApi.FrmPostVote)
	}
}
