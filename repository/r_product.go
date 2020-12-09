package repository

import (
	"apiapi/model"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

// ProductSet 注入
var ProductSet = wire.NewSet(wire.Struct(new(Product), "*"))

type Product struct {
	DB *gorm.DB
}

func (z *Product) FindAll() []model.Product {
	var products []model.Product
	z.DB.Find(&products)

	return products
}

func (z *Product) FindByID(id uint) model.Product {
	var product model.Product
	z.DB.First(&product, id)

	return product
}

func (z *Product) Save(product model.Product) model.Product {
	z.DB.Save(&product)

	return product
}

func (z *Product) Delete(product model.Product) {
	z.DB.Delete(&product)
}
