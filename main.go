package main

import (
	"BlueBell/dao/mysql"
	"BlueBell/dao/redis"
	"BlueBell/logger"
	"BlueBell/settings"
	"fmt"
)

func main() {
	if err := logger.Init(); err != nil {
		fmt.Println("init logger failed,err:", err)
		return
	}

	if err := settings.Init(); err != nil {
		fmt.Println("init setting failed,err:", err)
		return
	}

	if err := mysql.Init(); err != nil {
		fmt.Println("init setting failed,err:", err)
		return
	}

	if err := redis.Init(); err != nil {
		fmt.Println("init setting failed,err:", err)
		return
	}

}
