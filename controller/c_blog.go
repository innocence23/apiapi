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

// BlogSet 注入
var BlogSet = wire.NewSet(wire.Struct(new(Blog), "*"))

type Blog struct {
	BlogService *service.Blog
}

func (z *Blog) FindAll(c *gin.Context) {
	Blogs := z.BlogService.FindAll()

	c.JSON(http.StatusOK, gin.H{"Blogs": dto.ToBlogDTOs(Blogs)})
}

func (z *Blog) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	Blog := z.BlogService.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"Blog": dto.ToBlogDTO(Blog)})
}

func (z *Blog) Create(c *gin.Context) {
	var BlogDTO dto.BlogDTO
	err := c.BindJSON(&BlogDTO)
	if err != nil {
		c.Status(http.StatusBadRequest)
		logrus.Error(err)
		return
	}

	createdBlog := z.BlogService.Save(dto.ToBlog(BlogDTO))

	c.JSON(http.StatusOK, gin.H{"Blog": dto.ToBlogDTO(createdBlog)})
}

func (z *Blog) Update(c *gin.Context) {
	var BlogDTO dto.BlogDTO
	err := c.BindJSON(&BlogDTO)
	if err != nil {
		c.Status(http.StatusBadRequest)
		logrus.Error(err)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	Blog := z.BlogService.FindByID(uint(id))
	if Blog == (model.Blog{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	Blog.Title = BlogDTO.Title
	Blog.State = BlogDTO.State
	z.BlogService.Save(Blog)

	c.Status(http.StatusOK)
}

func (z *Blog) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	Blog := z.BlogService.FindByID(uint(id))
	if Blog == (model.Blog{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	z.BlogService.Delete(Blog)

	c.Status(http.StatusOK)
}
