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

func (p *ProductRepository) FindAll() []model.Product {
	var products []model.Product
	p.DB.Find(&products)

	return products
}

func (p *ProductRepository) FindByID(id uint) model.Product {
	var product model.Product
	p.DB.First(&product, id)

	return product
}

func (p *ProductRepository) Save(product model.Product) model.Product {
	p.DB.Save(&product)

	return product
}

func (p *ProductRepository) Delete(product model.Product) {
	p.DB.Delete(&product)
}
