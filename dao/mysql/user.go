package mysql

import (
	"BlueBell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

// 盐
const screty = "WANG_ZHANG"

// InsertUser 数据库插入一条新纪录
func InsertUser(user *models.User) error {
	//密码加密
	password := encryptPassword(user.Password)
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err := db.Exec(sqlStr, user.UserID, user.Username, password)
	if err != nil {
		return err
	}
	return nil
}

// CheckUserExist 查看指定用户名是否存在
func CheckUserExist(username string) error {
	sqlStr := `select count(*) from user where username=?`
	var count int
	err := db.Get(&count, sqlStr, username)
	if count > 0 {
		return ErrorUserExist
	}
	return err
}

// 加密密码
func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(screty))
	return hex.EncodeToString(h.Sum([]byte(password)))
}

func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id ,username, password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	password := encryptPassword(oPassword)
	if password != user.Password {
		err = ErrorUserNotExist
		return err
	}
	return
}

//根据id获取用户id

func GetUserByID(id uint64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id ,username from user where user_id = ?`
	err = db.Get(user, sqlStr, id)
	return
}
