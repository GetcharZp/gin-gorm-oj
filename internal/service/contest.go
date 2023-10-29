package service

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"getcharzp.cn/define"
	"getcharzp.cn/helper"
	"getcharzp.cn/models"
	"github.com/gin-gonic/gin"
)

// ContestCreate
// @Tags 管理员私有方法
// @Summary 竞赛创建
// @Accept json
// @Param authorization header string true "authorization"
// @Param data body define.ContestBasic true "contestBasic"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /admin/contest-create [post]
func ContestCreate(c *gin.Context) {
	in := new(define.ContestBasic)
	err := c.ShouldBindJSON(in)
	if err != nil {
		log.Println("[JsonBind Error] : ", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数错误",
		})
		return
	}

	if in.Name == "" || in.Content == "" || len(in.ProblemBasics) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不能为空",
		})
		return
	}

	identity := helper.GetUUID()
	data := &models.ContestBasic{
		Identity:  identity,
		Name:      in.Name,
		Content:   in.Content,
		StartAt:   models.MyTime(helper.ToTime(in.StartAt)),
		EndAt:     models.MyTime(helper.ToTime(in.EndAt)),
		CreatedAt: models.MyTime(time.Now()),
		UpdatedAt: models.MyTime(time.Now()),
	}

	contestProblems := make([]*models.ContestProblem, 0)
	for _, id := range in.ProblemBasics {
		contestProblems = append(contestProblems, &models.ContestProblem{
			ContestId: data.ID,
			ProblemId: uint(id),
			CreatedAt: models.MyTime(time.Now()),
			UpdatedAt: models.MyTime(time.Now()),
		})
	}
	data.ContestProblems = contestProblems

	// 创建竞赛
	err = models.DB.Create(data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Contest Create Error:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"identity": data.Identity,
		},
	})
}

// GetContestList
// @Tags 公共方法
// @Summary 竞赛列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /contest-list [get]
func GetContestList(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetContestList Page strconv Error:", err)
		return
	}
	page = (page - 1) * size
	var count int64
	keyword := c.Query("keyword")

	list := make([]*models.ContestBasic, 0)
	err = models.GetContestList(keyword).Distinct("`contest_basic`.`id`").Count(&count).Error
	if err != nil {
		log.Println("GetContestList Count Error:", err)
		return
	}
	err = models.GetContestList(keyword).Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Println("Get Contest List Error:", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})
}
