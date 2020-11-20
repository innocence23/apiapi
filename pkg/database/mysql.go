package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var DB *gorm.DB
var err error

func Init() {
	dns := fmt.Sprintf(
		"%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.dbname"),
	)

	DB, err = gorm.Open(viper.GetString("db.type"), dns)
	if err != nil {
		logrus.Panic("连接数据库失败", err)
	}

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	// 启用Logger，显示详细日志
	DB.LogMode(true)
}

func CloseDB() {
	err := DB.Close()
	if err != nil {
		logrus.Warn("数据库连接关闭失败")
	}
}
