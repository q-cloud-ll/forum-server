package initialize

import (
	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"

	"forum/global"
	"forum/utils"
)

func OtherInit() {
	dr, err := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.GVA_CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}
	err = InitWx()
	if err != nil {
		global.GVA_LOG.Error("InitWx failed,err:", zap.Error(err))
	}
	InitSnowflake()
	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)
}
