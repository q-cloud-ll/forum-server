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

type FollowApi struct{}

// FrmAddFollowUser 添加关注接口
func (fa *FollowApi) FrmAddFollowUser(c *gin.Context) {
	var a frmReq.AddFollower
	if err := c.ShouldBindJSON(&a); err != nil {
		global.GVA_LOG.Error("FrmAddFollowUser param with invalid, err:", zap.Error(err))
		response.FailWithMessage(xerr.REUQEST_PARAM_ERROR, c)
		return
	}

	followerId := utils.GetUserUuid(c).String()
	err := followService.FrmAddFollower(followerId, a.FolloweeId)
	if err != nil {
		global.GVA_LOG.Error("添加关注失败，err:", zap.Error(err))
		response.FailWithMessage(xerr.DB_ERROR, c)
		return
	}

	response.OkWithMessage("添加关注成功", c)
}

// FrmUnFollowUser 取消关注
func (fa *FollowApi) FrmUnFollowUser(c *gin.Context) {
	var a frmReq.AddFollower
	if err := c.ShouldBindJSON(&a); err != nil {
		global.GVA_LOG.Error("FrmUnFollowUser param with invalid, err:", zap.Error(err))
		response.FailWithMessage(xerr.REUQEST_PARAM_ERROR, c)
		return
	}

	followerId := utils.GetUserUuid(c).String()

	err := followService.FrmUnFollower(followerId, a.FolloweeId)
	if err != nil {
		global.GVA_LOG.Error("取消关注失败，err:", zap.Error(err))
		response.FailWithMessage(xerr.DB_ERROR, c)
		return
	}

	response.OkWithMessage("取消关注成功", c)
}

// FrmGetFollowers 获取关注列表
func (fa *FollowApi) FrmGetFollowers(c *gin.Context) {
	var a frmReq.GetFollowers
	if err := c.ShouldBindJSON(&a); err != nil {
		global.GVA_LOG.Error("FrmGetFollowers param with invalid, err:", zap.Error(err))
		response.FailWithMessage(xerr.REUQEST_PARAM_ERROR, c)
		return
	}

	data, total, err := followService.FrmGetFollowers(a)
	if err != nil {
		global.GVA_LOG.Error("获取关注列表失败，err:", zap.Error(err))
		response.FailWithMessage(xerr.DB_ERROR, c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     data,
		Page:     a.Page,
		PageSize: a.PageSize,
		Total:    total,
	}, "获取关注列表成功", c)
}

// FrmGetFollowees 获取粉丝列表
func (fa *FollowApi) FrmGetFollowees(c *gin.Context) {
	var a frmReq.GetFollowers
	if err := c.ShouldBindJSON(&a); err != nil {
		global.GVA_LOG.Error("FrmGetFollowees param with invalid, err:", zap.Error(err))
		response.FailWithMessage(xerr.REUQEST_PARAM_ERROR, c)
		return
	}

	data, total, err := followService.FrmGetFollowees(a)
	if err != nil {
		global.GVA_LOG.Error("获取粉丝列表失败，err:", zap.Error(err))
		response.FailWithMessage(xerr.DB_ERROR, c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     data,
		Page:     a.Page,
		PageSize: a.PageSize,
		Total:    total,
	}, "获取粉丝列表成功", c)
}
