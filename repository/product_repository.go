package repository

import (
	"apiapi/model"

	"github.com/jinzhu/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepostiory(DB *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: DB}
}

func (z *ProductRepository) FindAll() []model.Product {
	var products []model.Product
	z.DB.Find(&products)

	return products
}

func (z *ProductRepository) FindByID(id uint) model.Product {
	var product model.Product
	z.DB.First(&product, id)

	return product
}

func (z *ProductRepository) Save(product model.Product) model.Product {
	z.DB.Save(&product)

	return product
}

func (z *ProductRepository) Delete(product model.Product) {
	z.DB.Delete(&product)
}
