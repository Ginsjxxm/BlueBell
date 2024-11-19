package controller

import (
	"BlueBell/dao/mysql"
	"BlueBell/logic"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func CommunityHandler(c *gin.Context) {
	//查询所有社区community_id,community_name,以列表形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, data)
}

func CommunityDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		if errors.Is(err, mysql.ErrorInvalidID) {
			ResponseError(c, CodeInvalidByID)
			return
		}
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, data)
}
