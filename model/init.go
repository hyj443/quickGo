package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 包含了当前db连接的信息
var DB *gorm.DB

// DataBase 中间件，初始化mysql连接
func DataBase(dsn string) {
	// initialize a new db connection, need to import driver first
	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic(fmt.Sprintf("连接数据库出现异常：%v", err))
	}

	DB = db

	// 会在shell中log sql语句
	db.LogMode(true)

	// 从当前的连接中，获取go原生的*sql.DB，通用数据库对象
	sqlDB := db.DB()

	// 设置连接池sqlDB，sqlDB是包含多个数据库连接的连接池，有的是open有的是idle

	// 设置空闲连接池的最大连接数 即保持等待连接的连接数，避免操作过程中频繁获取连接/释放连接
	// 默认是有2个连接是一直保持的，不释放的，等待需要使用的用户使用
	sqlDB.SetMaxIdleConns(300)

	// 设置连接到数据库的最大连接数
	// 最多有500个并发打开数据库的连接
	sqlDB.SetMaxOpenConns(500)

	// 设置连接可以重用的最大周期，即超时时间，从创建开始算的
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	// 将构建的模型迁移为数据库的表
	migration()
}
