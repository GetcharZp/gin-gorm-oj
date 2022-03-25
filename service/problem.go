package service

import (
	"getcharzp.cn/define"
	"getcharzp.cn/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// GetProblemList
// @Tags 公共方法
// @Summary 问题列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Param category_identity query string false "category_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /problem-list [get]
func GetProblemList(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetProblemList Page strconv Error:", err)
		return
	}
	page = (page - 1) * size
	var count int64
	keyword := c.Query("keyword")
	categoryIdentity := c.Query("category_identity")

	list := make([]*models.Problem, 0)
	tx := models.GetProblemList(keyword, categoryIdentity)

	err = tx.Debug().Count(&count).Omit("content").Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Println("Get Problem List Error:", err)
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

// GetProblemDetail
// @Tags 公共方法
// @Summary 问题详情
// @Param identity query string true "问题的唯一标识"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Success 400 {string} json "{"code":"-1","msg":""}"
// @Router /problem-detail [get]
func GetProblemDetail(c *gin.Context) {
	problemIdentity := c.Query("identity")
	if problemIdentity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "问题标识必传",
		})
		return
	}
	data := new(models.Problem)
	err := models.DB.Model(new(models.Problem)).Where("identity = ?", problemIdentity).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "该问题不存在",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Gorm First models.Problem Error:" + err.Error(),
		})
		return
	}
	// 手动查询关联的分类信息
	categories := make([]*models.Category, 0)
	err = models.DB.Debug().Model(new(models.Category)).Where("id in ? ", strings.Split(data.CategoryId, ",")).Find(&categories).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Category Error:" + err.Error(),
		})
		return
	}
	data.Categories = categories
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}
