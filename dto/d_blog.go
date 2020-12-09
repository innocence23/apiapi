package dto

import "apiapi/model"

type BlogDTO struct {
	ID      int    `json:"id,omitempty"`
	TagID   int    `json:"tag_id,omitempty"` // 标签ID
	Title   string `json:"title,omitempty"`  // 文章标题
	Desc    string `json:"desc,omitempty"`   // 简述
	Content string `json:"content,omitempty"`
	State   int    `json:"state,omitempty"` // 状态 0为禁用1为启用
}

type BlogQueryParam struct {
	Title      string `json:"title,omitempty"`
	State      uint   `json:"state,omitempty"`
	Pagination `json:"pagination,omitempty"`
	OrderField `json:"order_field,omitempty"`
}

func ToBlog(BlogDTO BlogDTO) model.Blog {
	return model.Blog{Title: BlogDTO.Title, Desc: BlogDTO.Desc}
}

func ToBlogDTO(Blog model.Blog) BlogDTO {
	return BlogDTO{ID: Blog.ID, Title: Blog.Title, Desc: Blog.Desc}
}

func ToBlogDTOs(Blogs []model.Blog) []BlogDTO {
	Blogdtos := make([]BlogDTO, len(Blogs))

	for i, itm := range Blogs {
		Blogdtos[i] = ToBlogDTO(itm)
	}

	return Blogdtos
}
