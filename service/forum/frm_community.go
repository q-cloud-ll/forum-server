package forum

import (
	"forum/dao/mysql"
	"forum/global"
	"forum/model/forum/response"

	"go.uber.org/zap"
)

type CommunityService struct{}

// FrmGetCommunityInfo 获取社区信息服务
func (cs *CommunityService) FrmGetCommunityInfo(communityId int64) (data []response.FrmCommunityInfo, err error) {
	data, err = mysql.FrmGetCommunityInfo(communityId)
	if err != nil {
		global.GVA_LOG.Error("mysql.FrmGetCommunityInfo(communityId) failed",
			zap.Int64("communityId", communityId),
			zap.Error(err))
		return nil, err
	}
	return
}
