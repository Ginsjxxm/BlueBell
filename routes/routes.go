package routers

import (
	"BlueBell/controller"
	_ "BlueBell/docs"
	"BlueBell/logger"
	"BlueBell/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"net/http"
)

// SetupRouter 处理注册请求函数
func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")
	v1.POST("/SignUp", controller.SignUpHandler)
	v1.POST("/Login", controller.LoginHandler)
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetHandlerPost)
		v1.GET("/post", controller.GetPostListHandler)
		v1.POST("/vote", controller.PostVoteController)

		v1.GET("/post2", controller.GetPostListHandler2)
	}

	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	//test
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
