package controller

import "github.com/google/wire"

// RepostiorySet model注入
var ControllerSet = wire.NewSet(
	ProductSet,
	BlogSet,
)
