package mysql

import "errors"

var (
	ErrorUserExists      = errors.New("用户已存在")
	ErrorUserNotExists   = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
	ErrorInvalidID       = errors.New("无效的ID")
)
