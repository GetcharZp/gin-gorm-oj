package models

import (
	"gorm.io/gorm"
)

type CategoryBasic struct {
	ID        uint           `gorm:"primarykey;" json:"id"`
	CreatedAt MyTime         `json:"created_at"`
	UpdatedAt MyTime         `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;" json:"deleted_at"`
	Identity  string         `gorm:"column:identity;type:varchar(36);" json:"identity"` // 分类的唯一标识
	Name      string         `gorm:"column:name;type:varchar(100);" json:"name"`        // 分类名称
	ParentId  int            `gorm:"column:parent_id;type:int(11);" json:"parent_id"`   // 父级ID
}

func (table *CategoryBasic) TableName() string {
	return "category_basic"
}
