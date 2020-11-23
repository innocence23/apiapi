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

// func initDB() *gorm.DB{
// 	db, err := gorm.Open("mysql", os.Getenv("DB_URL"))
// 	if err != nil {
// 		panic(err)
// 	}

// 	db.AutoMigrate(&product.Product{})

// 	return db
// }

// func main() {
// 	db := initDB()
// 	defer db.Close()

// 	productAPI := InitProductAPI(db)

// 	r := gin.Default()

// 	r.GET("/products", productAPI.FindAll)
// 	r.GET("/products/:id", productAPI.FindByID)
// 	r.POST("/products", productAPI.Create)
// 	r.PUT("/products/:id", productAPI.Update)
// 	r.DELETE("/products/:id", productAPI.Delete)

// 	err := r.Run()
// 	if err != nil {
// 		panic(err)
// 	}
// }
