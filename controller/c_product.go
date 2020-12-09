package controller

import (
	"apiapi/dto"
	"apiapi/model"
	"apiapi/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

// ProductSet 注入
var ProductSet = wire.NewSet(wire.Struct(new(Product), "*"))

type Product struct {
	ProductService *service.Product
}

func (z *Product) FindAll(c *gin.Context) {
	products := z.ProductService.FindAll()

	c.JSON(http.StatusOK, gin.H{"products": dto.ToProductDTOs(products)})
}

func (z *Product) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := z.ProductService.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"product": dto.ToProductDTO(product)})
}

func (z *Product) Create(c *gin.Context) {
	var ProductDTO dto.ProductDTO
	err := c.BindJSON(&ProductDTO)
	if err != nil {
		c.Status(http.StatusBadRequest)
		logrus.Error(err)
		return
	}

	createdProduct := z.ProductService.Save(dto.ToProduct(ProductDTO))

	c.JSON(http.StatusOK, gin.H{"product": dto.ToProductDTO(createdProduct)})
}

func (z *Product) Update(c *gin.Context) {
	var ProductDTO dto.ProductDTO
	err := c.BindJSON(&ProductDTO)
	if err != nil {
		c.Status(http.StatusBadRequest)
		logrus.Error(err)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	product := z.ProductService.FindByID(uint(id))
	if product == (model.Product{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	product.Code = ProductDTO.Code
	product.Price = ProductDTO.Price
	z.ProductService.Save(product)

	c.Status(http.StatusOK)
}

func (z *Product) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := z.ProductService.FindByID(uint(id))
	if product == (model.Product{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	z.ProductService.Delete(product)

	c.Status(http.StatusOK)
}
