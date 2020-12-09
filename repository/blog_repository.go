package repository

import (
	"apiapi/model"

	"github.com/jinzhu/gorm"
)

type BlogRepository struct {
	DB *gorm.DB
}

func NewBlogRepostiory(DB *gorm.DB) *BlogRepository {
	return &BlogRepository{DB: DB}
}

func (z *BlogRepository) FindAll() []model.Blog {
	var Blogs []model.Blog
	z.DB.Find(&Blogs)

	return Blogs
}

func (z *BlogRepository) FindByID(id uint) model.Blog {
	var Blog model.Blog
	z.DB.First(&Blog, id)

	return Blog
}

func (z *BlogRepository) Save(Blog model.Blog) model.Blog {
	z.DB.Save(&Blog)

	return Blog
}

func (z *BlogRepository) Delete(Blog model.Blog) {
	z.DB.Delete(&Blog)
}
