package forum

import (
	v1 "forum-server/api/v1"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (f *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("user")
	//userRouterWithoutRecord := router.Group("user")
	frmUserApi := v1.ApiGroupApp.ForumApiGroup.UserApi
	{
		userRouter.POST("register", frmUserApi.FrmUserRegister)
		userRouter.POST("login", frmUserApi.FrmUserLogin)
	}
}
