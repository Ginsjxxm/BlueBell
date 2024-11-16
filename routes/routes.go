package routers

import (
	"BlueBell/controller"
	"BlueBell/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.POST("/SignUp", controller.SignUpHandler)
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
	})
	return r
}
