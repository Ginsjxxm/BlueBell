package controller

import (
	"BlueBell/dao/mysql"
	"BlueBell/logic"
	"BlueBell/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strings"
)

// 处理gin输入前缀
func removeTopStruct(fields map[string]string) map[string]string {
	res := make(map[string]string)
	for k, v := range fields {
		res[k[strings.Index(k, ".")+1:]] = v
	}
	return res
}

func SignUpHandler(c *gin.Context) {
	//参数校验
	p := models.ParamSignUp{}
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParams)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParams, removeTopStruct(errs.Translate(trans)))
		return
	}
	//业务处理
	err := logic.SignUp(&p)
	if err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServeBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	//获取参数
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidPassword, removeTopStruct(errs.Translate(trans)))
		return
	}

	//业务处理
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("Login with invalid param", zap.String("username", p.Username), zap.Error(err))
		ResponseError(c, CodeInvalidPassword)
		return
	}
	//返回响应
	ResponseSuccess(c, token)
}
