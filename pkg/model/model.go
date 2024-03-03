package model

import (
	"github.com/runningape/goblog/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {

	var err error

	config := mysql.New(mysql.Config{
		DSN: "root:dyhuangZz223@tcp(127.0.0.1:3306)/goblog?charset=utf8&parseTime=True&loc=local",
	})

	DB, err = gorm.Open(config, &gorm.Config{})

	logger.LogError(err)

	return DB
}
