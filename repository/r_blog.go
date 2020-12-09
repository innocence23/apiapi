package repository

import (
	"apiapi/model"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

// BlogSet 注入
var BlogSet = wire.NewSet(wire.Struct(new(Blog), "*"))

type Blog struct {
	DB *gorm.DB
}

func (z *Blog) FindAll() []model.Blog {
	var Blogs []model.Blog
	z.DB.Find(&Blogs)

	return Blogs
}

func (z *Blog) FindByID(id uint) model.Blog {
	var Blog model.Blog
	z.DB.First(&Blog, id)

	return Blog
}

func (z *Blog) Save(Blog model.Blog) model.Blog {
	z.DB.Save(&Blog)

	return Blog
}

func (z *Blog) Delete(Blog model.Blog) {
	z.DB.Delete(&Blog)
}
