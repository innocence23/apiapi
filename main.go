package main

import (
	"apiapi/pkg/cache"
	"apiapi/pkg/config"
	"apiapi/pkg/database"
	"apiapi/pkg/xlog"
	"apiapi/router"
)

func main() {
	//初始化配置
	config.Init()

	//初始化日志
	xlog.Init()

	//初始化db
	database.Init()

	//初始化redis
	cache.Init()

	//test queue nsq
	//queue.TestNsq()

	r := router.Init()
	r.Run(":8899") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
