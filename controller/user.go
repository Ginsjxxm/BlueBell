package controller

import (
	"BlueBell/logic"
	"BlueBell/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func SignUpHandler(c *gin.Context) {
	//参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	//业务处理
	logic.SignUp(p)
	//返回响应
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for k, v := range fields {
		res[k[strings.Index(k, ".")+1:]] = v
	}
	return res
}
