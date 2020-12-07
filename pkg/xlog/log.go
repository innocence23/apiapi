package xlog

import (
	"apiapi/pkg/config"
	"os"

	"github.com/sirupsen/logrus"
)

//Init log
func Init() {
	// 设置日志格式为json格式
	cfg := config.C

	if cfg.ENV == "production" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		file, err := os.OpenFile(cfg.Log.OutputFile, os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			logrus.SetOutput(file)
		} else {
			panic("日志文件创建失败")
		}
		// 设置日志级别为warn以上
		logrus.SetLevel(logrus.Level(cfg.Log.Level))

		//设置行号
		logrus.SetReportCaller(true)
	} else {
		// 设置日志级别为info以上
		logrus.SetLevel(logrus.Level(cfg.Log.TestLevel))
		logrus.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	}
}
