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
	"gorm.io/gorm"
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

// GetContestDetail
// @Tags 公共方法
// @Summary 竞赛详情
// @Param identity query string false "contest identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /contest-detail [get]
func GetContestDetail(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "竞赛唯一标识不能为空",
		})
		return
	}
	data := new(models.ContestBasic)
	err := models.DB.Where("identity = ?", identity).
		Preload("ContestProblems").Preload("ContestProblems.ProblemBasic").
		Preload("ContestUsers").
		First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "竞赛不存在",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Contest Detail Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}

// ContestModify
// @Tags 管理员私有方法
// @Summary 竞赛修改
// @Param authorization header string true "authorization"
// @Param data body define.ContestBasic true "ContestBasic"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /admin/contest-modify [put]
func ContestModify(c *gin.Context) {
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
	if in.Identity == "" || in.Name == "" || in.Content == "" || len(in.ProblemBasics) == 0 || in.StartAt == 0 || in.EndAt == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不能为空",
		})
		return
	}

	if err := models.DB.Transaction(func(tx *gorm.DB) error {
		// 竞赛基础信息保存 contest_basic
		contestBasic := &models.ContestBasic{
			Name:      in.Name,
			Content:   in.Content,
			StartAt:   models.MyTime(helper.ToTime(in.StartAt)),
			EndAt:     models.MyTime(helper.ToTime(in.EndAt)),
			UpdatedAt: models.MyTime(time.Now()),
		}
		err := tx.Where("identity = ?", in.Identity).Updates(contestBasic).Error
		if err != nil {
			return err
		}
		// 查询竞赛详情
		err = tx.Where("identity = ?", in.Identity).First(contestBasic).Error
		if err != nil {
			return err
		}

		// 关联问题题目的更新
		// 1、删除已存在的关联关系
		err = tx.Where("contest_id = ?", contestBasic.ID).Delete(new(models.ContestProblem)).Error
		if err != nil {
			return err
		}
		// 2、新增新的关联关系
		cps := make([]*models.ContestProblem, 0)
		for _, id := range in.ProblemBasics {
			cps = append(cps, &models.ContestProblem{
				ContestId: contestBasic.ID,
				ProblemId: uint(id),
				CreatedAt: models.MyTime(time.Now()),
				UpdatedAt: models.MyTime(time.Now()),
			})
		}
		err = tx.Create(&cps).Error
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Contest Modify Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "竞赛修改成功",
	})
}

// ContestDelete
// @Tags 管理员私有方法
// @Summary 竞赛删除
// @Param authorization header string true "authorization"
// @Param identity query string true "identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /admin/contest-delete [delete]
func ContestDelete(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不能为空",
		})
		return
	}
	if err := models.DB.Transaction(func(tx *gorm.DB) error {
		cbs := &models.ContestBasic{}
		// 查询竞赛是否存在
		err := tx.Where("identity = ?", identity).First(cbs).Error
		if err != nil {
			return err
		}

		// 删除问题的关联
		err = tx.Where("contest_id = ?", cbs.ID).Delete(new(models.ContestProblem)).Error
		if err != nil {
			return err
		}

		// 删除用户的关联
		err = tx.Where("contest_id = ?", cbs.ID).Delete(new(models.ContestUser)).Error
		if err != nil {
			return err
		}

		// 删除竞赛
		err = tx.Where("identity = ?", identity).Delete(cbs).Error
		return err
	}); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Contest Delete Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

// ContestRegistration
// @Tags 用户私有方法
// @Summary 竞赛报名
// @Param authorization header string true "authorization"
// @Param contest_identity query string true "contest_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/contest-registration [post]
func ContestRegistration(c *gin.Context) {
	contestIdentity := c.Query("contest_identity")
	if contestIdentity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不能为空",
		})
		return
	}
	// 查询竞赛是否存在
	cb := &models.ContestBasic{}
	err := models.DB.Where("identity = ?", contestIdentity).First(cb).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "竞赛不存在",
		})
		return
	}

	// 判断竞赛是否过期

	if time.Now().After(time.Time(cb.EndAt)) {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "竞赛已开始或已结束",
		})
		return
	}
	u, _ := c.Get("user_claims")
	userClaim := u.(*helper.UserClaims)

	// 判断是否报名
	err = models.DB.Where("contest_id = ? AND user_identity = ?", cb.ID, userClaim.Identity).First(&models.ContestUser{}).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	if err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "竞赛已报名",
		})
		return
	}

	// 进行报名
	cu := &models.ContestUser{
		ContestId:    cb.ID,
		UserIdentity: userClaim.Identity,
		CreatedAt:    models.MyTime(time.Now()),
		UpdatedAt:    models.MyTime(time.Now()),
	}
	err = models.DB.Create(cu).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常, 创建报名信息失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "报名成功",
	})
}
