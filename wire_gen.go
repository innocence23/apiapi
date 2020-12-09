// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"apiapi/API"
	"apiapi/app"
	"apiapi/repository"
	"apiapi/service"
	"github.com/jinzhu/gorm"
)

// Injectors from wire.go:

func initProductAPI(db *gorm.DB) *app.BllAPI {
	productRepository := repository.NewProductRepostiory(db)
	productService := service.NewProductService(productRepository)
	productAPI := API.NewProductAPI(productService)
	blogRepository := repository.NewBlogRepostiory(db)
	blogService := service.NewBlogService(blogRepository)
	blogAPI := API.NewBlogAPI(blogService)
	bllAPI := &app.BllAPI{
		ProductAPI: productAPI,
		BlogAPI:    blogAPI,
	}
	return bllAPI
}
