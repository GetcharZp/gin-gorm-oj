package models

import (
	"gorm.io/gorm"
)

type ContestBasic struct {
	ID              uint              `gorm:"primarykey;" json:"id"`
	CreatedAt       MyTime            `json:"created_at"`
	UpdatedAt       MyTime            `json:"updated_at"`
	DeletedAt       gorm.DeletedAt    `gorm:"index;" json:"deleted_at"`
	StartAt         MyTime            `json:"start_at"`
	EndAt           MyTime            `json:"end_at"`
	Identity        string            `gorm:"column:identity;type:varchar(36);" json:"identity"`            // 竞赛的唯一标识
	Name            string            `gorm:"column:name;type:varchar(100);" json:"name"`                   // 竞赛名称
	Content         string            `gorm:"column:content;type:text;" json:"content"`                     // 竞赛描述
	ContestProblems []*ContestProblem `gorm:"foreignKey:contest_id;references:id;" json:"contest_problems"` // 关联题目表
}

func (table *ContestBasic) TableName() string {
	return "contest_basic"
}
