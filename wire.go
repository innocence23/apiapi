// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"apiapi/API"
	"apiapi/app"
	"apiapi/repository"
	"apiapi/service"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func initProductAPI(db *gorm.DB) *app.BllAPI {
	wire.Build(
		repository.NewProductRepostiory, service.NewProductService, API.NewProductAPI,
		app.InjectorSet,
	)
	return new(app.BllAPI)
}
