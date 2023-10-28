package models

import (
	"gorm.io/gorm"
)

type ContestProblem struct {
	ID           uint           `gorm:"primarykey;" json:"id"`
	CreatedAt    MyTime         `json:"created_at"`
	UpdatedAt    MyTime         `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index;" json:"deleted_at"`
	ContestId    uint           `gorm:"column:contest_id;type:int(11);" json:"contest_id"`         // 问题的ID
	ProblemId    uint           `gorm:"column:problem_id;type:int(11);" json:"problem_id"`         // 分类的ID
	ProblemBasic *ProblemBasic  `gorm:"foreignKey:id;references:problem_id;" json:"problem_basic"` // 关联问题的基础信息表
}

func (table *ContestProblem) TableName() string {
	return "contest_problem"
}
