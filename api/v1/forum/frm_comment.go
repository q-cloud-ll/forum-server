package forum

import (
	"forum-server/global"
	"forum-server/model/common/response"
	frmReq "forum-server/model/forum/request"
	"forum-server/utils"
	"forum-server/utils/xerr"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CommentApi struct{}

// FrmCreateComment 创建用户评论接口
func (ca *CommentApi) FrmCreateComment(c *gin.Context) {
	var cr frmReq.Comment
	if err := c.ShouldBindQuery(&cr); err != nil {
		global.GVA_LOG.Error("FrmCreateComment param with invalid", zap.Error(err))
		response.FailWithMessage(xerr.REUQEST_PARAM_ERROR, c)
		return
	}
	err := utils.Verify(cr, utils.CreateComment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetUserUuid(c)
	if err := commentService.FrmCreateComment(&cr, uid); err != nil {
		global.GVA_LOG.Error("创建评论失败", zap.Error(err))
		response.FailWithMessage(xerr.DB_ERROR, c)
		return
	}

	response.OkWithMessage(xerr.OK, c)
}
