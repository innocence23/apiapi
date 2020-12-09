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

type BlogAPI struct {
	BlogService *service.BlogService
}

func NewBlogAPI(z *service.BlogService) *BlogAPI {
	return &BlogAPI{BlogService: z}
}

func (z *BlogAPI) FindAll(c *gin.Context) {
	Blogs := z.BlogService.FindAll()

	c.JSON(http.StatusOK, gin.H{"Blogs": dto.ToBlogDTOs(Blogs)})
}

func (z *BlogAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	Blog := z.BlogService.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"Blog": dto.ToBlogDTO(Blog)})
}

func (z *BlogAPI) Create(c *gin.Context) {
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

func (z *BlogAPI) Update(c *gin.Context) {
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

	Blog.Code = BlogDTO.Code
	Blog.Price = BlogDTO.Price
	z.BlogService.Save(Blog)

	c.Status(http.StatusOK)
}

func (z *BlogAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	Blog := z.BlogService.FindByID(uint(id))
	if Blog == (model.Blog{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	z.BlogService.Delete(Blog)

	c.Status(http.StatusOK)
}
