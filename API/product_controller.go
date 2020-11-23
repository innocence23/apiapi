package API

import (
	. "apiapi/dto"
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

func NewProductAPI(p *service.ProductService) *ProductAPI {
	return &ProductAPI{ProductService: p}
}

func (p *ProductAPI) FindAll(c *gin.Context) {
	products := p.ProductService.FindAll()

	c.JSON(http.StatusOK, gin.H{"products": ToProductDTOs(products)})
}

func (p *ProductAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"product": ToProductDTO(product)})
}

func (p *ProductAPI) Create(c *gin.Context) {
	var productDTO ProductDTO
	err := c.BindJSON(&productDTO)
	if err != nil {
		c.Status(http.StatusBadRequest)
		logrus.Error(err)
		return
	}

	createdProduct := p.ProductService.Save(ToProduct(productDTO))

	c.JSON(http.StatusOK, gin.H{"product": ToProductDTO(createdProduct)})
}

func (p *ProductAPI) Update(c *gin.Context) {
	var productDTO ProductDTO
	err := c.BindJSON(&productDTO)
	if err != nil {
		c.Status(http.StatusBadRequest)
		logrus.Error(err)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.FindByID(uint(id))
	if product == (model.Product{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	product.Code = productDTO.Code
	product.Price = productDTO.Price
	p.ProductService.Save(product)

	c.Status(http.StatusOK)
}

func (p *ProductAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.FindByID(uint(id))
	if product == (model.Product{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.ProductService.Delete(product)

	c.Status(http.StatusOK)
}
