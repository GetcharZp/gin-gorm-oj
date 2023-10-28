package service

import (
	"log"
	"net/http"
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
