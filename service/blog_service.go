package service

import (
	"apiapi/model"
	"apiapi/repository"
)

type BlogService struct {
	BlogRepository *repository.BlogRepository
}

func NewBlogService(z *repository.BlogRepository) *BlogService {
	return &BlogService{BlogRepository: z}
}

func (z *BlogService) FindAll() []model.Blog {
	return z.BlogRepository.FindAll()
}

func (z *BlogService) FindByID(id uint) model.Blog {
	return z.BlogRepository.FindByID(id)
}

func (z *BlogService) Save(Blog model.Blog) model.Blog {
	z.BlogRepository.Save(Blog)

	return Blog
}

func (z *BlogService) Delete(Blog model.Blog) {
	z.BlogRepository.Delete(Blog)
}
