package main

import (
	"apiapi/API"
	"apiapi/repository"
	"apiapi/service"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func initProductAPI(db *gorm.DB) *API.ProductAPI {
	wire.Build(repository.NewProductRepostiory, service.NewProductService, API.NewProductAPI)
	return &API.ProductAPI{}
}
