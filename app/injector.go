package app

import (
	"apiapi/API"

	"github.com/google/wire"
)

// InjectorSet 注入Injector
//var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))
var InjectorSet = wire.NewSet(wire.Struct(new(BllAPI), "*"))


type BllAPI struct {
	ProductAPI *API.ProductAPI
}
