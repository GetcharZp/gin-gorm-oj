package router

import (
	"getcharzp.cn/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	// 路由规则
	r.GET("/ping", service.Ping)
	r.GET("/problem-list", service.GetProblemList)

	return r
}
