package forum

import (
	"forum-server/global"
	"forum-server/model/common/response"
	"forum-server/model/forum"
	FrmUserReq "forum-server/model/forum/request"
	systemRes "forum-server/model/forum/response"
	systemReq "forum-server/model/system/request"
	"forum-server/utils"
	"forum-server/utils/xerr"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserApi struct{}

// FrmUserRegister 用户注册接口
func (userApi *UserApi) FrmUserRegister(c *gin.Context) {
	var register FrmUserReq.Register
	if err := c.ShouldBindJSON(&register); err != nil {
		global.GVA_LOG.Error("FrmUserRegister param with invalid", zap.Error(err))
		response.FailWithMessage(xerr.REUQEST_PARAM_ERROR, c)
		return
	}

	if err := userService.FrmUserRegister(&register); err != nil {
		global.GVA_LOG.Error("用户注册失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("注册成功！", c)
}

// FrmUserLogin 用户登录接口
func (userApi *UserApi) FrmUserLogin(c *gin.Context) {
	var login FrmUserReq.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		global.GVA_LOG.Error("FrmUserLogin param with invalid", zap.Error(err))
		response.FailWithMessage(xerr.REUQEST_PARAM_ERROR, c)
		return
	}
	user, err := userService.FrmUserLogin(&login)
	if err != nil {
		global.GVA_LOG.Error("用户名不存在或密码错误", zap.Error(err))
		response.FailWithMessage("用户名不存在或者密码错误", c)
		return
	}
	if user.Enable != 1 {
		global.GVA_LOG.Error("登录失败，该用户被禁止登录")
		response.FailWithMessage("用户被禁止登录", c)
		return
	}
	userApi.TokenNext(c, *user)
	return
}

// TokenNext 将用户信息保存进token
func (userApi *UserApi) TokenNext(c *gin.Context, user forum.FrmUser) {
	j := utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)}
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:     user.UserId,
		ID:       user.ID,
		NickName: user.Nickname,
		Username: user.Username,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}
	response.OkWithDetailed(systemRes.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", c)
}
