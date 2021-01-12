package app

import (
	"apiapi/controller"

	"github.com/google/wire"
)

// InjectorSet 注入Injector
//var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))
var InjectorSet = wire.NewSet(wire.Struct(new(BllAPI), "*"))

type BllAPI struct {
	ProductAPI *controller.Product
	BlogAPI    *controller.Blog
	TagAPI     *controller.Tag
}
