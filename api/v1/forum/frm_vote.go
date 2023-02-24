package forum

import (
	"forum/global"
	"forum/model/common/response"
	frmReq "forum/model/forum/request"
	"forum/utils"
	"forum/utils/xerr"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type VoteApi struct{}

// FrmPostVote 帖子投票接口
func (voteApi *VoteApi) FrmPostVote(c *gin.Context) {
	var v frmReq.FrmVoteData
	if err := c.ShouldBindJSON(&v); err != nil {
		global.GVA_LOG.Error("FrmPostVote param with invalid, err:", zap.Error(err))
		response.FailWithMessage(xerr.REUQEST_PARAM_ERROR, c)
		return
	}
	if err := utils.Verify(v, utils.VotePostVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 获取当前请求的用户id
	userId := utils.GetUserUuid(c)
	if err := voteService.FrmVotePost(userId, &v); err != nil {
		global.GVA_LOG.Error("投票失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("投票成功", c)
}
