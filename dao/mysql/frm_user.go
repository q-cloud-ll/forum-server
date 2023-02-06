package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"forum-server/global"
	"forum-server/model/forum"
	"forum-server/utils"
)

const secret = "forum-server:qll"

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

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))

	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
