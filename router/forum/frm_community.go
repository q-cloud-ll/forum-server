package forum

import (
	v1 "forum/api/v1"

	"github.com/gin-gonic/gin"
)

type CommunityRouter struct{}

// InitCommunityRouterPublic 公共路由
func (r *CommunityRouter) InitCommunityRouterPublic(router *gin.RouterGroup) {
	communityRouter := router.Group("community")
	frmCommunityApi := v1.ApiGroupApp.ForumApiGroup.CommunityApi
	{
		communityRouter.GET("getCommunityInfo", frmCommunityApi.FrmGetCommunityInfo)
	}
}
