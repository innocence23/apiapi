package model

type Tag struct {
	ID         int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Name       string `gorm:"column:name" json:"name"`                         // 标签名称
	CreatedOn  int    `gorm:"column:created_on;default:0" json:"created_on"`   // 创建时间
	CreatedBy  string `gorm:"column:created_by" json:"created_by"`             // 创建人
	ModifiedOn int    `gorm:"column:modified_on;default:0" json:"modified_on"` // 修改时间
	ModifiedBy string `gorm:"column:modified_by" json:"modified_by"`           // 修改人
	DeletedOn  int    `gorm:"column:deleted_on;default:0" json:"deleted_on"`
	State      int    `gorm:"column:state;default:1" json:"state"` // 状态 0为禁用、1为启用
}
