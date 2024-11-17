package routers

import (
	"BlueBell/controller"
	"BlueBell/logger"
	"BlueBell/middlewares"
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
	r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	//没有注册的路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "not found",
		})
	})
	return r
}
