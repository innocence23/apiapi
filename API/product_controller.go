package API

import (
	"apiapi/dto"
	"apiapi/model"
	"apiapi/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ProductAPI struct {
	ProductService *service.ProductService
}

func NewProductAPI(z *service.ProductService) *ProductAPI {
	return &ProductAPI{ProductService: z}
}

func (z *ProductAPI) FindAll(c *gin.Context) {
	products := z.ProductService.FindAll()

	c.JSON(http.StatusOK, gin.H{"products": dto.ToProductDTOs(products)})
}

func (z *ProductAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := z.ProductService.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"product": dto.ToProductDTO(product)})
}

func (z *ProductAPI) Create(c *gin.Context) {
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

func (z *ProductAPI) Update(c *gin.Context) {
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

func (z *ProductAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := z.ProductService.FindByID(uint(id))
	if product == (model.Product{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	z.ProductService.Delete(product)

	c.Status(http.StatusOK)
}
