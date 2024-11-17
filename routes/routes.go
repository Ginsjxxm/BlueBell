package routers

import (
	"BlueBell/controller"
	"BlueBell/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetupRouter 处理注册请求函数
func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.POST("/SignUp", controller.SignUpHandler)
	r.POST("/Login", controller.LoginHandler)
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
	})
	return r
}
