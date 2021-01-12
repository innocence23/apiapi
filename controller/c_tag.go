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

// TagSet 注入
var TagSet = wire.NewSet(wire.Struct(new(Tag), "*"))

type Tag struct {
	TagService *service.Tag
}

func (z *Tag) FindAll(c *gin.Context) {
	Tags := z.TagService.FindAll()

	c.JSON(http.StatusOK, gin.H{"Tags": dto.ToTagDTOs(Tags)})
}

func (z *Tag) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	Tag := z.TagService.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"Tag": dto.ToTagDTO(Tag)})
}

func (z *Tag) Create(c *gin.Context) {
	var TagDTO dto.TagDTO
	err := c.BindJSON(&TagDTO)
	if err != nil {
		c.Status(http.StatusBadRequest)
		logrus.Error(err)
		return
	}

	createdTag := z.TagService.Save(dto.ToTag(TagDTO))

	c.JSON(http.StatusOK, gin.H{"Tag": dto.ToTagDTO(createdTag)})
}

func (z *Tag) Update(c *gin.Context) {
	var TagDTO dto.TagDTO
	err := c.BindJSON(&TagDTO)
	if err != nil {
		c.Status(http.StatusBadRequest)
		logrus.Error(err)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	Tag := z.TagService.FindByID(uint(id))
	if Tag == (model.Tag{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	Tag.Name = TagDTO.Name
	Tag.State = TagDTO.State
	z.TagService.Save(Tag)

	c.Status(http.StatusOK)
}

func (z *Tag) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	Tag := z.TagService.FindByID(uint(id))
	if Tag == (model.Tag{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	z.TagService.Delete(Tag)

	c.Status(http.StatusOK)
}
