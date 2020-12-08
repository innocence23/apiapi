package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type ProductQueryParam struct {
	Code       string `json:"code,omitempty"`
	Price      uint   `json:"price,omitempty"`
	Pagination `json:"pagination,omitempty"`
	OrderField `json:"order_field,omitempty"`
}
