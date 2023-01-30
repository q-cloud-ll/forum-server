package forum

import (
	v1 "forum-server/api/v1"

	"github.com/gin-gonic/gin"
)

type PostRouter struct{}

// InitPostRouter 过验证的路由
func (p *PostRouter) InitPostRouter(router *gin.RouterGroup) {
	postRouter := router.Group("post")
	frmPostApi := v1.ApiGroupApp.ForumApiGroup.PostApi
	{
		postRouter.POST("createPost", frmPostApi.FrmPostCreatePost)
	}
}

func (p *PostRouter) InitPostRouterPublic(router *gin.RouterGroup) {
	postRouter := router.Group("post")
	frmPostApi := v1.ApiGroupApp.ForumApiGroup.PostApi
	{
		postRouter.GET("postList", frmPostApi.FrmPostGetPostList)
	}
}
