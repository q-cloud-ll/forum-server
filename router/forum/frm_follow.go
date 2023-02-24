package forum

import (
	v1 "forum/api/v1"

	"github.com/gin-gonic/gin"
)

type FollowRouter struct{}

// InitFollowRouter 过验证的路由
func (p *FollowRouter) InitFollowRouter(router *gin.RouterGroup) {
	followRouter := router.Group("follow")
	frmFollowApi := v1.ApiGroupApp.ForumApiGroup.FollowApi
	{
		followRouter.POST("addFollowUser", frmFollowApi.FrmAddFollowUser)
		followRouter.DELETE("followUser", frmFollowApi.FrmUnFollowUser)
		followRouter.POST("getFollowerList", frmFollowApi.FrmGetFollowers)
		followRouter.POST("getFolloweeList", frmFollowApi.FrmGetFollowees)
	}
}
