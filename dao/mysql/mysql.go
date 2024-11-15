package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.Get("mysql.user"),
		viper.Get("mysql.password"),
		viper.Get("mysql.host"),
		viper.Get("mysql.port"),
		viper.Get("mysql.dbname"),
	)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("connect database error:", err)
		return
	}

	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle"))
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open"))

	return
}
