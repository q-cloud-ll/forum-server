package forum

import (
	v1 "forum/api/v1"
	"forum/global"

	"github.com/gin-gonic/gin"
)

type QRCodeRouter struct{}

func (qrr *QRCodeRouter) InitQRCodeRouterPublic(router *gin.RouterGroup) {
	wxRouter := router.Group("wechat")
	frmWeChatApi := v1.ApiGroupApp.ForumApiGroup.WeChatApi
	{
		wxRouter.GET("getTicket", frmWeChatApi.FrmGetQRCodeTicket)
		wxRouter.GET("wx", func(c *gin.Context) {
			server := global.GVA_WX.GetServer(c.Request, c.Writer)
			err := server.Serve()
			if err != nil {
				return
			}
			server.Send()
		})
		//qrcodeRouter.GET("wx", frmWeChatApi.FrmWxCheckSignature)
		wxRouter.POST("wx", frmWeChatApi.FrmWxGetTicket)
		wxRouter.GET("wx/sign", frmWeChatApi.FrmSignInfo)
		//qrcodeRouter.POST("scanLogin", frmWeChatApi.FrmScanLogin)
	}
}
