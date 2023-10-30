package models

import (
	"gorm.io/gorm"
)

type ContestUser struct {
	ID           uint           `gorm:"primarykey;" json:"id"`
	CreatedAt    MyTime         `json:"created_at"`
	UpdatedAt    MyTime         `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index;" json:"deleted_at"`
	ContestId    uint           `gorm:"column:contest_id;type:int(11);" json:"contest_id"`               // 问题的ID
	UserIdentity string         `gorm:"column:user_identity;type:int(11);" json:"user_identity"`         // 分类的ID
	UserBasic    *UserBasic     `gorm:"foreignKey:identity;references:user_identity;" json:"user_basic"` // 关联问题的基础信息表
}

func (table *ContestUser) TableName() string {
	return "contest_user"
}
