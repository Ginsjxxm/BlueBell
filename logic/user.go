package logic

import (
	"BlueBell/dao/mysql"
	"BlueBell/models"
	"BlueBell/pkg/snowflake"
	"errors"
)

var ()

func SignUp(p *models.ParamSignUp) (err error) {
	//用户名是否存在
	err = mysql.CheckUserExist(p.Username)
	if errors.Is(err, mysql.ErrorUserExist) {
		return err
	}
	if err != nil {
		return err
	}
	//生成UID
	userID, _ := snowflake.GenID()

	//构建user实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	//存入数据库
	err = mysql.InsertUser(user)
	if err != nil {
		return err
	}
	return nil
}

func Login(p *models.ParamLogin) (err error) {
	user := models.User{
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.Login(&user)
}
