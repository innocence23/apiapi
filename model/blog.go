package model

type Blog struct {
	ID         int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	TagID      int    `gorm:"column:tag_id;default:0" json:"tag_id"` // 标签ID
	Title      string `gorm:"column:title" json:"title"`             // 文章标题
	Desc       string `gorm:"column:desc" json:"desc"`               // 简述
	Content    string `gorm:"column:content" json:"content"`
	CreatedOn  int    `gorm:"column:created_on" json:"created_on"`
	CreatedBy  string `gorm:"column:created_by" json:"created_by"`             // 创建人
	ModifiedOn int    `gorm:"column:modified_on;default:0" json:"modified_on"` // 修改时间
	ModifiedBy string `gorm:"column:modified_by" json:"modified_by"`           // 修改人
	DeletedOn  int    `gorm:"column:deleted_on;default:0" json:"deleted_on"`
	State      int    `gorm:"column:state;default:1" json:"state"` // 状态 0为禁用1为启用
}
