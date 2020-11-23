package router

import (
	"apiapi/pkg/database"

	"github.com/gin-gonic/gin"
)

// Init router
func Init() *gin.Engine {
	r := gin.Default()

	productAPI := initProductAPI(database.DB)
	r.GET("/products", productAPI.FindAll)
	r.GET("/products/:id", productAPI.FindByID)
	r.POST("/products", productAPI.Create)
	r.PUT("/products/:id", productAPI.Update)
	r.DELETE("/products/:id", productAPI.Delete)

	return r
}
