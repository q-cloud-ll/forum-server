package redis

import (
	"forum-server/global"
	"time"
)

// FrmSetAccessToken 存放AccessToken
func FrmSetAccessToken(token string) error {
	return global.GVA_REDIS.Set(ctx, getRedisKey(KeyQRCodeTicket), token, 300*time.Second).Err()
}

// FrmGetAccessToken 获取AccessToken
func FrmGetAccessToken() (string, error) {
	return global.GVA_REDIS.Get(ctx, getRedisKey(KeyQRCodeTicket)).Result()
}

// FrmSetWxSceneStr 设置二维码请求scene字段
func FrmSetWxSceneStr(str string) error {
	return global.GVA_REDIS.Set(ctx, getRedisKey(str), "", 300*time.Second).Err()

}

// FrmGetWxSceneStr 获取二维码请求scene字段
func FrmGetWxSceneStr(str string) (string, error) {
	return global.GVA_REDIS.Get(ctx, getRedisKey(str)).Result()
}

// FrmSetEvent 公众号消息回复
func FrmSetEvent(eventKey, fromUserName string) (string, error) {
	return global.GVA_REDIS.Set(ctx, eventKey, fromUserName, 300*time.Second).Result()
}
