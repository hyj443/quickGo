package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// DB 包含了当前db连接的信息
var DB *gorm.DB

// DataBase 中间件，初始化mysql连接
func DataBase(dsn string) {

	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic(fmt.Sprintf("连接数据库出现异常：%v", err))
	}

	DB = db

	// true代表detailed logs，false代表no log，default, will only print error logs
	// 会在shell中log sql语句
	db.LogMode(true)

	sqlDB := db.DB()

	sqlDB.SetMaxIdleConns(50)

	sqlDB.SetMaxOpenConns(100)

	sqlDB.SetConnMaxLifetime(time.Second * 30)

	// 将构建的模型迁移为数据库的表
	migration()

}
