package router

import (
	_ "getcharzp.cn/docs"
	"getcharzp.cn/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	// Swagger 配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 路由规则

	// 问题
	r.GET("/problem-list", service.GetProblemList)
	r.GET("/problem-detail", service.GetProblemDetail)

	// 用户
	r.GET("/user-detail", service.GetUserDetail)
	r.POST("/login", service.Login)

	// 提交记录
	r.GET("/submit-list", service.GetSubmitList)

	return r
}
