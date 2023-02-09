package forum

import (
	"forum-server/global"
	"forum-server/model/common/response"
	"forum-server/utils/xerr"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type QRCodeApi struct{}

func (qra *QRCodeApi) FrmGetQRCode(c *gin.Context) {
	qrCodeInfo, err := qrcodeService.FrmGenerateQRCode(c)
	if err != nil {
		global.GVA_LOG.Error("获取二维码失败", zap.Error(err))
		response.FailWithMessage(xerr.QRCodeGetFailErr, c)
		return
	}
	response.OkWithData(qrCodeInfo, c)
}

func (qra *QRCodeApi) FrmAskQRCode(c *gin.Context) {

}

func (qra *QRCodeApi) FrmScanLogin(c *gin.Context) {

}
