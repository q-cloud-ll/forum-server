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

type CommentApi struct{}

// FrmCreateComment 创建用户评论接口
func (ca *CommentApi) FrmCreateComment(c *gin.Context) {
	var pc frmReq.PostComment
	if err := c.ShouldBindJSON(&pc); err != nil {
		global.GVA_LOG.Error("FrmCreateComment param with invalid,err:", zap.Error(err))
		response.FailWithMessage(xerr.REUQEST_PARAM_ERROR, c)
		return
	}
	err := utils.Verify(pc, utils.CreateComment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetUserUuid(c).String()
	if err := commentService.FrmCreateComment(&pc, uid); err != nil {
		global.GVA_LOG.Error("创建评论失败", zap.Error(err))
		response.FailWithMessage(xerr.DB_ERROR, c)
		return
	}

	response.OkWithMessage(xerr.OK, c)
}

// FrmGetPostCommentList 获取帖子详情的评论列表
func (ca *CommentApi) FrmGetPostCommentList(c *gin.Context) {
	p := &frmReq.CommentList{
		Page:     1,
		PageSize: 10,
		Order:    frmReq.OrderScore,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		global.GVA_LOG.Error("FrmGetPostComment with invalid query,err:", zap.Error(err))
		response.FailWithMessage(xerr.REUQEST_PARAM_ERROR, c)
		return
	}
	data, total, err := commentService.FrmGetPostCommentList(p)
	if err != nil {
		global.GVA_LOG.Error("获取评论失败，err:", zap.Error(err))
		response.FailWithMessage(xerr.DB_ERROR, c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     data,
		Page:     int(p.Page),
		PageSize: int(p.PageSize),
		Total:    int64(total),
	}, "获取评论成功", c)
}

// FrmGetChildrenCommentList 获取子评论接口
func (ca *CommentApi) FrmGetChildrenCommentList(c *gin.Context) {
	p := &frmReq.CdCommentList{
		Page:     1,
		PageSize: 10,
		Order:    frmReq.OrderScore,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		global.GVA_LOG.Error("FrmGetChildrenCommentList with invalid query,err:", zap.Error(err))
		response.FailWithMessage(xerr.REUQEST_PARAM_ERROR, c)
		return
	}
	data, total, err := commentService.FrmGetChildrenCommentList(p)
	if err != nil {
		global.GVA_LOG.Error("获取子评论失败，err:", zap.Error(err))
		response.FailWithMessage(xerr.DB_ERROR, c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     data,
		Page:     int(p.Page),
		PageSize: int(p.PageSize),
		Total:    int64(total),
	}, "获取子评论成功", c)
}

// FrmStarComment 给评论点赞接口
func (ca *CommentApi) FrmStarComment(c *gin.Context) {
	var sc frmReq.StarComment
	if err := c.ShouldBindJSON(&sc); err != nil {
		global.GVA_LOG.Error("FrmStarComment params with invalid,err:", zap.Error(err))
		response.FailWithMessage(xerr.REUQEST_PARAM_ERROR, c)
		return
	}
	uid := utils.GetUserUuid(c).String()
	err := commentService.FrmStarComment(&sc, uid)
	if err != nil {
		global.GVA_LOG.Error("评论点赞失败，err:", zap.Error(err))
		response.FailWithMessage(xerr.DB_ERROR, c)
		return
	}
	response.OkWithMessage("评论点赞成功", c)
}
