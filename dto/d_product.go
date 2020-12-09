package dto

import "apiapi/model"

type ProductDTO struct {
	ID    uint   `json:"id,string,omitempty"`
	Code  string `json:"code"`
	Price uint   `json:"price,string"`
}

type ProductQueryParam struct {
	Code       string `json:"code,omitempty"`
	Price      uint   `json:"price,omitempty"`
	Pagination `json:"pagination,omitempty"`
	OrderField `json:"order_field,omitempty"`
}

func ToProduct(productDTO ProductDTO) model.Product {
	return model.Product{Code: productDTO.Code, Price: productDTO.Price}
}

func ToProductDTO(product model.Product) ProductDTO {
	return ProductDTO{ID: product.ID, Code: product.Code, Price: product.Price}
}

func ToProductDTOs(products []model.Product) []ProductDTO {
	productdtos := make([]ProductDTO, len(products))

	for i, itm := range products {
		productdtos[i] = ToProductDTO(itm)
	}

	return productdtos
}
