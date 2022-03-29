# 基于Gin、Gorm 实现的在线练习系统

> 往期视频：https://www.bilibili.com/video/BV1ha4y1E76t?spm_id_from=333.999.0.0
> 语言：PHP、MVC框架：TP6、中间件：Redis
> 
> 语言：Golang、框架：Gin、GORM

GOLANG下载网址： https://golang.google.cn/dl/

参考文档 Module：https://www.kancloud.cn/aceld/golang/1958311

GORM中文官网：https://gorm.io/zh_CN/docs/

GIN中文官网：https://gin-gonic.com/zh-cn/docs/

## 整合Swagger
参考文档： https://github.com/swaggo/gin-swagger
接口访问地址：http://localhost:8080/swagger/index.html
```text
// GetProblemList
// @Tags 公共方法
// @Summary 问题列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /problem-list [get]
```

## 安装 jwt
```shell
go get github.com/dgrijalva/jwt-go
```

## 配置 
+ 将 MailPassword 配置到环境变量中