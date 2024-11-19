package logic

import (
	"BlueBell/dao/mysql"
	"BlueBell/models"
	"BlueBell/pkg/jwt"
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

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	//传递指针可以拿到userid
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	//生成jwt
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return nil, err
	}
	user.Token = token
	return user, nil
}
