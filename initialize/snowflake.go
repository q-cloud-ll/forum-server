package initialize

import (
	"forum-server/global"
	"forum-server/utils"

	"go.uber.org/zap"
)

func InitSnowflake() {
	snowflakeCfg := global.GVA_CONFIG.Snowflake
	if err := utils.Init(snowflakeCfg.StartTime, snowflakeCfg.MachineID); err != nil {
		global.GVA_LOG.Error("init snowflake failed, err:", zap.Error(err))
		return
	}
}
