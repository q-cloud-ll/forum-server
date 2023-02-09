package forum

import (
	"fmt"
	"forum-server/dao/redis"
	"forum-server/global"
	"forum-server/model/forum"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type QRCodeService struct{}

// FrmGenerateQRCode 生成二维码
func (qrs *QRCodeService) FrmGenerateQRCode(c *gin.Context) (qrInfo *forum.FrmQRCodeInfo, err error) {
	// 获取uuid
	mTicket := uuid.NewV4().String()
	qrCodeUrl := fmt.Sprintf("http://127.0.0.1:8889/api/qrcode/scanLogin?mticket=%s", mTicket)

	authInfo := &forum.AuthInfo{
		Token:  "",
		Status: 0,
	}
	err = redis.FrmGenQRCode(mTicket, authInfo)
	if err != nil {
		global.GVA_LOG.Error("redis.FrmGenQRCode save qrcode info failed, err:", zap.Error(err))
		return nil, err
	}

	c.SetCookie("m_ticket", mTicket, 3*60, "/", "localhost", true, true)

	qrCodeInfo := &forum.FrmQRCodeInfo{
		CodeUrl: qrCodeUrl,
	}

	return qrCodeInfo, err
}

//func (qrs *QRCodeService) FrmScanLogin(c *gin.Context) (err error) {
//	mTicket := c.Query("m_ticket")
//	if mTicket == "" {
//		global.GVA_LOG.Error("c.Query('m_ticket') failed")
//		return errors.New(xerr.QRCodeRetryErr)
//	}
//
//	res, err := redis.FrmScanLogin(mTicket)
//	if err != nil {
//		global.GVA_LOG.Error("redis.FrmScanLogin(mTicket) get failed, err:", zap.Error(err))
//		return errors.New(xerr.QRCodeRetryErr)
//	}
//
//}
