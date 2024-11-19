package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const ContextUserIDKey = "userid"

var ErrorNotLogin = errors.New("用户未登录")

// GetCurrentUserID 根据中间件寻找key
func GetCurrentUserID(c *gin.Context) (userId uint64, err error) {
	uid, ok := c.Get(ContextUserIDKey)
	if !ok {
		err = ErrorNotLogin
		return
	}
	userId, ok = uid.(uint64)
	if !ok {
		err = ErrorNotLogin
		return
	}
	return
}
