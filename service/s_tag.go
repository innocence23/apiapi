package service

import (
	"apiapi/model"
	"apiapi/repository"

	"github.com/google/wire"
)

// TagSet 注入
var TagSet = wire.NewSet(wire.Struct(new(Tag), "*"))

type Tag struct {
	TagRepository *repository.Tag
}

func (z *Tag) FindAll() []model.Tag {
	return z.TagRepository.FindAll()
}

func (z *Tag) FindByID(id uint) model.Tag {
	return z.TagRepository.FindByID(id)
}

func (z *Tag) Save(Tag model.Tag) model.Tag {
	z.TagRepository.Save(Tag)

	return Tag
}

func (z *Tag) Delete(Tag model.Tag) {
	z.TagRepository.Delete(Tag)
}
