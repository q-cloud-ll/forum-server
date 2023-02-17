package forum

import (
	"forum-server/global"
	"forum-server/model/common/response"
	"forum-server/model/forum"
	frmReq "forum-server/model/forum/request"
	"forum-server/utils"
	"forum-server/utils/xerr"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type PostApi struct{}

// FrmPostCreatePost 创建帖子接口
func (postApi *PostApi) FrmPostCreatePost(c *gin.Context) {
	var p forum.FrmPost
	if err := c.ShouldBindJSON(&p); err != nil {
		global.GVA_LOG.Error("FrmPostCreatePost param with invalid, err:", zap.Error(err))
		response.FailWithMessage(xerr.REUQEST_PARAM_ERROR, c)
		return
	}
	err := utils.Verify(p, utils.CreatePostVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 从token获取uuid，从而知道这个帖子的作者
	userId := utils.GetUserUuid(c)
	p.AuthorId = userId
	if err := postService.FrmPostCreatePost(&p); err != nil {
		global.GVA_LOG.Error("创建帖子失败", zap.Error(err))
		response.FailWithMessage(xerr.DB_ERROR, c)
		return
	}
	response.OkWithMessage(xerr.OK, c)
}

// FrmPostGetPostList 获取帖子列表接口
func (postApi *PostApi) FrmPostGetPostList(c *gin.Context) {
	p := &frmReq.PostList{
		Page:     1,
		PageSize: 10,
		Order:    frmReq.OrderScore,
	}
	// 参数校验
	if err := c.ShouldBindQuery(p); err != nil {
		global.GVA_LOG.Error("FrmPostGetPostList with invalid query,err:", zap.Error(err))
		response.FailWithMessage(xerr.REUQEST_PARAM_ERROR, c)
		return
	}
	// 将请求参数传入获取数据
	list, total, err := postService.FrmPostGetPostList(p)
	if err != nil {
		global.GVA_LOG.Error("获取帖子列表失败", zap.Error(err))
		response.FailWithMessage(xerr.DB_ERROR, c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     int(p.Page),
		PageSize: int(p.PageSize),
	}, "查询成功", c)
}

// FrmPostDetail 获取帖子详情接口
func (postApi *PostApi) FrmPostDetail(c *gin.Context) {
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		global.GVA_LOG.Error("FrmPostDetail with param invalid", zap.Error(err))
		response.FailWithMessage(xerr.REUQEST_PARAM_ERROR, c)
		return
	}
	data, err := postService.FrmPostDetail(pid)
	if err != nil {
		global.GVA_LOG.Error("获取帖子详情失败", zap.Error(err))
		response.FailWithMessage(xerr.DB_ERROR, c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}
