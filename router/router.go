package router

import (
	"github.com/gin-gonic/gin"
	"sky-take-out-go/controller/admin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//  新增员工路由
	router.POST("/admin/employee", admin.Save)
	router.GET("/admin/employee/page", admin.Page)
	router.POST("/admin/employee/login", admin.Login)
	
	return router
}
