package forum

import (
	"errors"
	"fmt"
	"forum/dao/mysql"
	"forum/global"
	"forum/model/forum"
	FrmUserReq "forum/model/forum/request"
	"forum/utils"
	"forum/utils/tool"

	uuid "github.com/satori/go.uuid"

	"go.uber.org/zap"
)

type UserService struct{}

// FrmUserRegister 用户注册服务
func (userService *UserService) FrmUserRegister(f *FrmUserReq.Register) (err error) {
	if err := mysql.CheckUserExist(f.Username); err != nil {
		global.GVA_LOG.Error("user already exists, err:", zap.Error(err))
		return err
	}
	userId := uuid.NewV4()
	nickName := "新用户_" + tool.Krand(8, tool.KC_RAND_KIND_ALL)
	user := &forum.FrmUser{
		UserId:   userId,
		Username: f.Username,
		Mobile:   f.Mobile,
		Password: f.Password,
		Nickname: nickName,
	}

	return mysql.InsertUser(user)
}

// FrmUserLogin 用户登录服务
func (userService *UserService) FrmUserLogin(f *FrmUserReq.Login) (userInter *forum.FrmUser, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}
	var user forum.FrmUser
	err = global.GVA_DB.Table("frm_users").Where("username = ?", f.Username).First(&user).Error
	fmt.Println(f.Password, user.Password, utils.BcryptCheck(f.Password, user.Password))
	if err == nil {
		if ok := utils.BcryptCheck(f.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &user, err
}
