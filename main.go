package main

import (
	"BlueBell/controller"
	"BlueBell/dao/mysql"
	"BlueBell/dao/redis"
	"BlueBell/logger"
	"BlueBell/pkg/snowflake"
	routers "BlueBell/routes"
	"BlueBell/settings"
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Bell
// @version 1.0
// @description BlueBell
// @termsOfService http://swagger.io/terms/

// @contact.name Admira.
// @contact.url http://www.swagger.io/support
// @contact.email 3288449152@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8081/
// @BasePath api/v1
func main() {
	//加载配置
	if err := settings.Init(); err != nil {
		fmt.Println("init setting failed,err:", err)
		return
	}
	//加载日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Println("init logger failed,err:", err)
		return
	}
	zap.L().Info("logger init success")
	//异步加载入内
	defer zap.L().Sync()
	//加载数据库
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Println("init setting failed,err:", err)
		return
	}
	defer mysql.Close()

	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Println("init redis failed,err:", err)
		return
	}
	defer redis.Close()

	//雪花id
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Println("init snowflake failed,err:", err)
		return
	}
	//validator验证
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Println("init trans failed,err:", err)
		return
	}

	//注册路由
	r := routers.SetupRouter(settings.Conf.Mode)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown Server ...")

	// 创建上下文，设置超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // 设置超时时间为 5 秒
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("Server Shutdown failed", zap.Error(err))
	}
	zap.L().Info("Server exiting")

}
