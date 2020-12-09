package service

import (
	"apiapi/model"
	"apiapi/repository"

	"github.com/google/wire"
)

// BlogSet 注入
var BlogSet = wire.NewSet(wire.Struct(new(Blog), "*"))

type Blog struct {
	BlogRepository *repository.Blog
}

func (z *Blog) FindAll() []model.Blog {
	return z.BlogRepository.FindAll()
}

func (z *Blog) FindByID(id uint) model.Blog {
	return z.BlogRepository.FindByID(id)
}

func (z *Blog) Save(Blog model.Blog) model.Blog {
	z.BlogRepository.Save(Blog)

	return Blog
}

func (z *Blog) Delete(Blog model.Blog) {
	z.BlogRepository.Delete(Blog)
}
