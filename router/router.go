package router

import (
	"apiapi/app"

	"github.com/gin-gonic/gin"
)

// Init router
func Init(a *app.BllAPI) *gin.Engine {
	r := gin.Default()

	r.GET("/products", a.ProductAPI.FindAll)
	r.GET("/products/:id", a.ProductAPI.FindByID)
	r.POST("/products", a.ProductAPI.Create)
	r.PUT("/products/:id", a.ProductAPI.Update)
	r.DELETE("/products/:id", a.ProductAPI.Delete)

	r.GET("/blogs", a.BlogAPI.FindAll)
	r.GET("/blogs/:id", a.BlogAPI.FindByID)
	r.POST("/blogs", a.BlogAPI.Create)
	r.PUT("/blogs/:id", a.BlogAPI.Update)
	r.DELETE("/blogs/:id", a.BlogAPI.Delete)

	r.GET("/tags", a.TagAPI.FindAll)
	r.GET("/tags/:id", a.TagAPI.FindByID)
	r.POST("/tags", a.TagAPI.Create)
	r.PUT("/tags/:id", a.TagAPI.Update)
	r.DELETE("/tags/:id", a.TagAPI.Delete)

	return r
}
