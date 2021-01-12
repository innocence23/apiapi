package dto

import "apiapi/model"

type TagDTO struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`        // 标签名称
	CreatedOn  int    `json:"created_on,omitempty"`  // 创建时间
	CreatedBy  string `json:"created_by,omitempty"`  // 创建人
	ModifiedOn int    `json:"modified_on,omitempty"` // 修改时间
	ModifiedBy string `json:"modified_by,omitempty"` // 修改人
	DeletedOn  int    `json:"deleted_on,omitempty"`
	State      int    `json:"state,omitempty"` // 状态 0为禁用、1为启用
}

type TagQueryParam struct {
	Title      string `json:"title,omitempty"`
	State      uint   `json:"state,omitempty"`
	Pagination `json:"pagination,omitempty"`
	OrderField `json:"order_field,omitempty"`
}

func ToTag(TagDTO TagDTO) model.Tag {
	return model.Tag{Name: TagDTO.Name, State: TagDTO.State}
}

func ToTagDTO(Tag model.Tag) TagDTO {
	return TagDTO{ID: Tag.ID, Name: Tag.Name, State: Tag.State}
}

func ToTagDTOs(Tags []model.Tag) []TagDTO {
	Tagdtos := make([]TagDTO, len(Tags))

	for i, itm := range Tags {
		Tagdtos[i] = ToTagDTO(itm)
	}

	return Tagdtos
}
