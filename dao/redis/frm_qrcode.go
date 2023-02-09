package redis

import (
	"forum-server/global"
	"forum-server/model/forum"
	"time"
)

// FrmGenQRCode 保存二维码信息
func FrmGenQRCode(mTicket string, authInfo *forum.AuthInfo) (err error) {
	_, err = global.GVA_REDIS.Set(ctx, getQRTicketKey(mTicket), authInfo, 3*time.Minute).Result()
	return
}

// FrmScanLogin 获取二维码信息
func FrmScanLogin(mTicket string) (res string, err error) {
	res, err = global.GVA_REDIS.Get(ctx, getQRTicketKey(mTicket)).Result()
	return
}
