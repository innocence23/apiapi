package service

import "github.com/google/wire"

// ServiceSet model注入
var ServiceSet = wire.NewSet(
	ProductSet,
	BlogSet,
	TagSet,
)
