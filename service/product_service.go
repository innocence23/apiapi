package service

import (
	"apiapi/model"
	"apiapi/repository"
)

type ProductService struct {
	ProductRepository *repository.ProductRepository
}

func NewProductService(z *repository.ProductRepository) *ProductService {
	return &ProductService{ProductRepository: z}
}

func (z *ProductService) FindAll() []model.Product {
	return z.ProductRepository.FindAll()
}

func (z *ProductService) FindByID(id uint) model.Product {
	return z.ProductRepository.FindByID(id)
}

func (z *ProductService) Save(product model.Product) model.Product {
	z.ProductRepository.Save(product)

	return product
}

func (z *ProductService) Delete(product model.Product) {
	z.ProductRepository.Delete(product)
}
