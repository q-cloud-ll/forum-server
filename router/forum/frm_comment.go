package forum

import (
	v1 "forum-server/api/v1"

	"github.com/gin-gonic/gin"
)

type CommentRouter struct{}

// InitCommentRouter 过验证的路由
func (p *CommentRouter) InitCommentRouter(router *gin.RouterGroup) {
	commentRouter := router.Group("comment")
	frmCommentApi := v1.ApiGroupApp.ForumApiGroup.CommentApi
	{
		commentRouter.POST("createComment", frmCommentApi.FrmCreateComment)
		commentRouter.POST("starComment", frmCommentApi.FrmStarComment)
		commentRouter.GET("getCommentList", frmCommentApi.FrmGetPostCommentList)
	}
}

// InitCommentRouterPublic 公共路由
func (p *CommentRouter) InitCommentRouterPublic(router *gin.RouterGroup) {
	commentRouter := router.Group("comment")
	frmCommentApi := v1.ApiGroupApp.ForumApiGroup.CommentApi
	{
		commentRouter.GET("getCdCommentList", frmCommentApi.FrmGetChildrenCommentList)
	}
}
