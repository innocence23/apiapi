package database

import (
	"apiapi/pkg/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

var DB *gorm.DB
var err error

func Init() {
	cfg := config.C.DB
	DB, err = gorm.Open(cfg.Type, cfg.DSN())
	if err != nil {
		logrus.Panic("连接数据库失败", err)
	}

	DB.DB().SetMaxIdleConns(cfg.MaxIdleConns)
	DB.DB().SetMaxOpenConns(cfg.MaxOpenConns)

	// 启用Logger，显示详细日志
	DB.LogMode(cfg.Debug)
}

func CloseDB() {
	err := DB.Close()
	if err != nil {
		logrus.Warn("数据库连接关闭失败")
	}
}
