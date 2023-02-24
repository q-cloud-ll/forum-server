package mysql

import (
	"fmt"
	"forum/global"
	"forum/model/forum"
	"forum/utils"

	uuid "github.com/satori/go.uuid"
)

//const secret = "forum:qll"

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	var count int64
	if err = global.GVA_DB.Table("frm_users").
		Where("username=?", username).
		Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExists
	}

	return
}

func InsertUser(u *forum.FrmUser) (err error) {
	u.Password = utils.BcryptHash(u.Password)
	err = global.GVA_DB.Table("frm_users").Create(&u).Error

	return err
}

//
//func encryptPassword(oPassword string) string {
//	h := md5.New()
//	h.Write([]byte(secret))
//
//	return hex.EncodeToString(h.Sum([]byte(oPassword)))
//}

// FrmGetUserById 根据uuid查询用户信息
func FrmGetUserById(uuid uuid.UUID) (user *forum.FrmUser, err error) {
	var u *forum.FrmUser
	err = global.GVA_DB.Table("frm_users").
		Select("user_id, nickname, avatar").
		Where("user_id = ?", uuid).
		Find(&u).
		Error
	return u, err
}

// FrmGetUserInfoById 根据uid查询用户信息
func FrmGetUserInfoById(uid string) (user *forum.FrmUserInfo, err error) {
	var u *forum.FrmUserInfo
	err = global.GVA_DB.Table("frm_users").
		Select("nickname, avatar").
		Where("user_id = ?", uid).
		Find(&u).
		Error
	return u, err
}

// FrmGetUserInfoByCommentId 根据uid查询用户信息
func FrmGetUserInfoByCommentId(commentId int64) (user *forum.FrmUserInfo, err error) {
	var u *forum.FrmUserInfo
	err = global.GVA_DB.Table("frm_comments fc").
		Select("fu.nickname as nickname, fu.avatar as avatar, fu.user_id as user_id").
		Joins("join frm_users fu on fu.user_id = fc.user_id").
		Where("fc.comment_id = ?", commentId).
		Find(&u).
		Error
	fmt.Println(u.Nickname, u.Avatar)
	return u, err
}
