package models

import (
	"gorm.io/gorm"
)

type ProblemBasic struct {
	ID                uint               `gorm:"primarykey;" json:"id"`
	CreatedAt         MyTime             `json:"created_at"`
	UpdatedAt         MyTime             `json:"updated_at"`
	DeletedAt         gorm.DeletedAt     `gorm:"index;" json:"deleted_at"`
	Identity          string             `gorm:"column:identity;type:varchar(36);" json:"identity"`                  // 问题表的唯一标识
	ProblemCategories []*ProblemCategory `gorm:"foreignKey:problem_id;references:id" json:"problem_categories"`      // 关联问题分类表
	Title             string             `gorm:"column:title;type:varchar(255);" json:"title"`                       // 文章标题
	Content           string             `gorm:"column:content;type:text;" json:"content"`                           // 文章正文
	MaxRuntime        int                `gorm:"column:max_runtime;type:int(11);" json:"max_runtime"`                // 最大运行时长
	MaxMem            int                `gorm:"column:max_mem;type:int(11);" json:"max_mem"`                        // 最大运行内存
	TestCases         []*TestCase        `gorm:"foreignKey:problem_identity;references:identity;" json:"test_cases"` // 关联测试用例表
	PassNum           int64              `gorm:"column:pass_num;type:int(11);" json:"pass_num"`                      // 通过次数
	SubmitNum         int64              `gorm:"column:submit_num;type:int(11);" json:"submit_num"`                  // 提交次数
}

func (table *ProblemBasic) TableName() string {
	return "problem_basic"
}

func GetProblemList(keyword, categoryIdentity string) *gorm.DB {
	tx := DB.Model(new(ProblemBasic)).Distinct("`problem_basic`.`id`").Select("DISTINCT(`problem_basic`.`id`), `problem_basic`.`identity`, "+
		"`problem_basic`.`title`, `problem_basic`.`max_runtime`, `problem_basic`.`max_mem`, `problem_basic`.`pass_num`, "+
		"`problem_basic`.`submit_num`, `problem_basic`.`created_at`, `problem_basic`.`updated_at`, `problem_basic`.`deleted_at` ").Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").
		Where("title like ? OR content like ? ", "%"+keyword+"%", "%"+keyword+"%")
	if categoryIdentity != "" {
		tx.Joins("RIGHT JOIN problem_category pc on pc.problem_id = problem_basic.id").
			Where("pc.category_id = (SELECT cb.id FROM category_basic cb WHERE cb.identity = ? )", categoryIdentity)
	}
	return tx.Order("problem_basic.id DESC")
}
