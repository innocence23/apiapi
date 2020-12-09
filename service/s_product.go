package service

import (
	"apiapi/model"
	"apiapi/repository"

	"github.com/google/wire"
)

// ProductSet 注入
var ProductSet = wire.NewSet(wire.Struct(new(Product), "*"))

type Product struct {
	ProductRepository *repository.Product
}

func (z *Product) FindAll() []model.Product {
	return z.ProductRepository.FindAll()
}

func (z *Product) FindByID(id uint) model.Product {
	return z.ProductRepository.FindByID(id)
}

func (z *Product) Save(product model.Product) model.Product {
	z.ProductRepository.Save(product)

	return product
}

func (z *Product) Delete(product model.Product) {
	z.ProductRepository.Delete(product)
}
