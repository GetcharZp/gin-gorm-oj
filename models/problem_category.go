package models

import "gorm.io/gorm"

type ProblemCategory struct {
	gorm.Model
	ProblemId  string    `gorm:"column:problem_id;type:varchar(36);" json:"problem_id"`   // 问题的唯一标识
	CategoryId string    `gorm:"column:category_id;type:varchar(36);" json:"category_id"` // 分类的唯一标识
	Category   *Category `gorm:"foreignKey:id;references:category_id"`
}

func (table *ProblemCategory) TableName() string {
	return "problem_category"
}
