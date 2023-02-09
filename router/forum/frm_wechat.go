package forum

import (
	v1 "forum-server/api/v1"

	"github.com/gin-gonic/gin"
)

type QRCodeRouter struct{}

func (qrr *QRCodeRouter) InitQRCodeRouterPublic(router *gin.RouterGroup) {
	qrcodeRouter := router.Group("wechat")
	frmQRCodeApi := v1.ApiGroupApp.ForumApiGroup.QRCodeApi
	{
		qrcodeRouter.GET("get", frmQRCodeApi.FrmGetQRCode)
		qrcodeRouter.GET("ask", frmQRCodeApi.FrmAskQRCode)
		qrcodeRouter.POST("scanLogin", frmQRCodeApi.FrmScanLogin)
	}
}
