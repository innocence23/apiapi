package repository

import (
	"apiapi/model"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

// TagSet 注入
var TagSet = wire.NewSet(wire.Struct(new(Tag), "*"))

type Tag struct {
	DB *gorm.DB
}

func (z *Tag) FindAll() []model.Tag {
	var Tags []model.Tag
	z.DB.Find(&Tags)

	return Tags
}

func (z *Tag) FindByID(id uint) model.Tag {
	var Tag model.Tag
	z.DB.First(&Tag, id)

	return Tag
}

func (z *Tag) Save(Tag model.Tag) model.Tag {
	z.DB.Save(&Tag)

	return Tag
}

func (z *Tag) Delete(Tag model.Tag) {
	z.DB.Delete(&Tag)
}
