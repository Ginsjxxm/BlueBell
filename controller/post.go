package controller

import (
	"BlueBell/logic"
	"BlueBell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// CreatePostHandler 创建帖子
func CreatePostHandler(c *gin.Context) {
	//获取参数
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	//回去请求帖子的用户id
	userID, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	//创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}

	//返回响应
	ResponseSuccess(c, nil)

}

// GetHandlerPost 获取帖子
func GetHandlerPost(c *gin.Context) {
	//获取url_id
	Opid := c.Param("id")
	pid, err := strconv.ParseInt(Opid, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
	}
	data, err := logic.GetPostByID(pid)
	if err != nil {
		zap.L().Error("logic.GetPostByID failed", zap.Error(err))
		c.JSON(
			http.StatusOK,
			gin.H{
				"msg": err.Error(),
			})
		return
	}

	ResponseSuccess(c, data)
}

// GetPostListHandler 分页请求帖子
func GetPostListHandler(c *gin.Context) {
	//获取分页参数
	limit, offset := getInfo(c)
	data, err := logic.GetPostList(offset, limit)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, data)
}

// GetPostListHandler2 升级版帖子列表
// 按照时间或者创建时间获取
func GetPostListHandler2(c *gin.Context) {
	//获取分页参数
	p := &models.ParamPostList{
		Limit:  0,
		Offset: 5,
		Order:  models.OrderTime,
	}
	err := c.ShouldBindQuery(p)
	if err != nil {
		zap.L().Error("logic.GetPostList2 failed", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	data, err := logic.GetPostListNew(p)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, data)

}

func getInfo(c *gin.Context) (limit int64, offset int64) {
	offsetStr := c.Query("offset")
	limitStr := c.Query("limit")
	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = 0
	}
	limit, err = strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		limit = 5
	}
	return limit, offset
}
