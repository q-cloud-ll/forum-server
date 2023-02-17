package forum

import (
	"forum-server/global"
	"forum-server/model/common/response"
	"forum-server/utils/xerr"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type CommunityApi struct{}

// FrmGetCommunityInfo 获取社区信息接口
func (ca *CommunityApi) FrmGetCommunityInfo(c *gin.Context) {
	cm := c.Query("community_id")
	if err := c.ShouldBindQuery(&cm); err != nil {
		global.GVA_LOG.Error("FrmGetCommunityInfo param with invalid, err:", zap.Error(err))
		response.FailWithMessage(xerr.REUQEST_PARAM_ERROR, c)
		return
	}

	communityId, _ := strconv.ParseInt(cm, 10, 64)
	data, err := communityService.FrmGetCommunityInfo(communityId)
	if err != nil {
		global.GVA_LOG.Error("获取社区信息失败", zap.Error(err))
		response.FailWithMessage("获取社区信息失败", c)
		return
	}

	response.OkWithDetailed(data, "获取社区信息成功", c)
}
