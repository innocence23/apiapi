package config

import (
	"fmt"

	"github.com/spf13/viper"
)

//Init config
func Init() {
	viper.SetConfigName("config") //  设置配置文件名 (不带后缀)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("config") // 比如添加当前目录
	//viper.AddConfigPath("/workspace/appName1") // 可以多次调用添加路径
	err := viper.ReadInConfig() // 搜索路径，并读取配置数据
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}
}
