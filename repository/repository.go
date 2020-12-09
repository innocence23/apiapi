package repository

import "github.com/google/wire"

// RepostiorySet model注入
var RepostiorySet = wire.NewSet(
	ProductSet,
	BlogSet,
)
