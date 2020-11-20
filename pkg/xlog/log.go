package xlog

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//Init log
func Init() {
	// 设置日志格式为json格式
	if viper.GetString("env") == "production" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		file, err := os.OpenFile(viper.GetString("log.logger_file"), os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			logrus.SetOutput(file)
		} else {
			panic("日志文件创建失败")
		}
		// 设置日志级别为warn以上
		logrus.SetLevel(logrus.WarnLevel)

		//设置行号
		logrus.SetReportCaller(true)
	} else {
		// 设置日志级别为info以上
		logrus.SetLevel(logrus.InfoLevel)
		logrus.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	}
}
