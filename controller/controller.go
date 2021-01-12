package controller

import "github.com/google/wire"

// ControllerSet model注入
var ControllerSet = wire.NewSet(
	ProductSet,
	BlogSet,
	TagSet,
)
