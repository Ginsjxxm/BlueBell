package logic

import (
	"BlueBell/dao/mysql"
	"BlueBell/models"
	"BlueBell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) {
	//用户名是否存在
	mysql.QueryUserByUsername()
	//生成UID
	snowflake.GenID()
	//存入数据库
	mysql.InsertUser()
}
