package models

import (
	"gorm.io/gorm"
)

type Problem struct {
	gorm.Model
	Identity   string `gorm:"column:identity;type:varchar(36);" json:"identity"`        // 问题表的唯一标识
	CategoryId string `gorm:"column:category_id;type:varchar(255);" json:"category_id"` // 分类ID，已逗号分隔
	Title      string `gorm:"column:title;type:varchar(255);" json:"title"`             // 文章标题
	Content    string `gorm:"column:content;type:text;" json:"content"`                 // 文章正文
	MaxRuntime int    `gorm:"column:max_runtime;type:int(11);" json:"max_runtime"`      // 最大运行时长
	MaxMem     int    `gorm:"column:max_mem;type:int(11);" json:"max_mem"`              // 最大运行内存
}

func (table *Problem) TableName() string {
	return "problem"
}

func GetProblemList(keyword string) *gorm.DB {
	return DB.Model(new(Problem)).
		Where("title like ? OR content like ? ", "%"+keyword+"%", "%"+keyword+"%")
}
