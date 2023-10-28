package define

import (
	"os"
)

var (
	DefaultPage = "1"
	DefaultSize = "20"
)

var MailPassword = os.Getenv("MailPassword")
var MysqlDNS = os.Getenv("MysqlDNS")

type ProblemBasic struct {
	Identity          string      `json:"identity"`           // 问题表的唯一标识
	Title             string      `json:"title"`              // 问题标题
	Content           string      `json:"content"`            // 问题内容
	ProblemCategories []int       `json:"problem_categories"` // 关联问题分类表
	MaxRuntime        int         `json:"max_runtime"`        // 最大运行时长
	MaxMem            int         `json:"max_mem"`            // 最大运行内存
	TestCases         []*TestCase `json:"test_cases"`         // 关联测试用例表
}

type TestCase struct {
	Input  string `json:"input"`  // 输入
	Output string `json:"output"` // 输出
}

type ContestBasic struct {
	Identity      string `json:"identity"`      // 竞赛的唯一标识
	Name          string `json:"name"`          // 竞赛名称
	Content       string `json:"content"`       // 竞赛描述
	ProblemBasics []int  `json:"problem_basic"` // 关联题目表id
	StartAt       int64  `json:"start_at"`      // 竞赛开启时间
	EndAt         int64  `json:"end_at"`        // 竞赛关闭时间
}

var (
	DateLayout            = "2006-01-02 15:04:05"
	ValidGolangPackageMap = map[string]struct{}{
		"bytes":   {},
		"fmt":     {},
		"math":    {},
		"sort":    {},
		"strings": {},
	}
)
