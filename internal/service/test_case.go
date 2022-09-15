package service

import (
	"getcharzp.cn/define"
	"getcharzp.cn/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetTestCase
// @Tags 管理员私有方法
// @Summary 测试案例列表
// @Param authorization header string true "authorization"
// @Param identity query string true "问题唯一标识"
// @Param page query int false "page"
// @Param size query int false "size"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /admin/test-case [get]
func GetTestCase(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetProblemList Page strconv Error:", err)
		return
	}
	page = (page - 1) * size
	problemIdentity := c.Query("identity")
	if problemIdentity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "问题唯一标识不能为空",
		})
		return
	}
	var count int64

	data := make([]*models.TestCase, 0)
	tx := models.DB.Model(new(models.TestCase)).Where("problem_identity = ?", problemIdentity).Count(&count)
	if c.Query("size") != "" {
		tx.Offset(page).Limit(page)
	}
	err = tx.Find(&data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get TestCase List Error:" + err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  data,
			"count": count,
		},
	})
}
