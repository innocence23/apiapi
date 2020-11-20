package service

import (
	"apiapi/model"
	"apiapi/repository"
)

type ProductService struct {
	ProductRepository *repository.ProductRepository
}

func NewProductService(p *repository.ProductRepository) *ProductService {
	return &ProductService{ProductRepository: p}
}

func (p *ProductService) FindAll() []model.Product {
	return p.ProductRepository.FindAll()
}

func (p *ProductService) FindByID(id uint) model.Product {
	return p.ProductRepository.FindByID(id)
}

func (p *ProductService) Save(product model.Product) model.Product {
	p.ProductRepository.Save(product)

	return product
}

func (p *ProductService) Delete(product model.Product) {
	p.ProductRepository.Delete(product)
}
