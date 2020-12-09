// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"apiapi/app"
	"apiapi/controller"
	"apiapi/repository"
	"apiapi/service"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func initProductAPI(db *gorm.DB) *app.BllAPI {
	wire.Build(
		repository.RepostiorySet,
		service.ServiceSet,
		controller.ControllerSet,
		app.InjectorSet,
	)
	return new(app.BllAPI)
}
