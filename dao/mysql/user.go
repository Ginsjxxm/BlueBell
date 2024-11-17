package mysql

import (
	"BlueBell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

var (
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或者密码错误")
	ErrorUserExist       = errors.New("用户已存在")
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

func Login(p *models.User) (err error) {
	oPassword := p.Password
	password := p.Password
	sqlStr := `select password from user where username=?`
	err = db.Get(&password, sqlStr, p.Username)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	oPassword = encryptPassword(oPassword)
	if password != oPassword {
		return ErrorInvalidPassword
	}
	return
}
