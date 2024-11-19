package mysql

import "errors"

var (
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或者密码错误")
	ErrorUserExist       = errors.New("用户已存在")
	ErrorInvalidID       = errors.New("无效的ID")
)
