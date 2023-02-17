package initialize

import (
	"forum-server/global"

	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
)

func InitWx() error {
	wc := wechat.NewWechat()
	//这里本地内存保存access_token，也可选择redis，memcache或者自定cache
	memory := cache.NewMemory()
	//cache.NewRedis(context.Background(), global.GVA_REDIS)
	wxCfg := &offConfig.Config{
		AppID:     global.GVA_CONFIG.WeChat.AppID,
		AppSecret: global.GVA_CONFIG.WeChat.AppSecret,
		Token:     global.GVA_CONFIG.WeChat.Token,
		//EncodingAESKey: "xxxx",
		Cache: memory,
	}
	//获取微信实例
	global.GVA_WX = wc.GetOfficialAccount(wxCfg)
	return nil
}
