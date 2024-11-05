package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sky-take-out-go/controller/admin/employee"
	"sky-take-out-go/model/entity"
	"sky-take-out-go/utils"
	"sky-take-out-go/controller/admin/category"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	
	// 员工管理路由
	{
		router.POST("/admin/employee", JwtHandler(), employee.Save)	//  新增员工路由
		router.POST("/admin/employee/login", employee.Login)	//  员工登录路由
		router.GET("/admin/employee/:id", JwtHandler(), employee.GetById)		// 根据ID查询员工路由	
		router.POST("/admin/employee/status/:status", JwtHandler(), employee.StartOrStop)		// 启用或停用员工路由
		router.PUT("/admin/employee/editPassword", JwtHandler(), employee.EditPassword)	 // 修改密码路由
		router.POST("/admin/employee/logout", JwtHandler(), employee.EmpLogout)		// 退出登录路由
		router.PUT("/admin/employee", JwtHandler(), employee.Edit)		// 修改员工信息路由
		router.GET("/admin/employee/page", JwtHandler(), employee.Page)	//  分页查询员工路由

	}
	
	// 分类管理路由
	{
		router.POST("/admin/category", JwtHandler(), category.Save)	//  新增分类路由
		router.GET("admin/category/page", JwtHandler(), category.Page)	//  分页查询分类路由
	}
	return router
}

func JwtHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("token")
		if token == "" {
			context.JSON(http.StatusInternalServerError, entity.Response{
				Code: 0,
				Data: nil,
				Msg:  nil,
			})
			context.Abort()
			return
		}
		claims, err := utils.ParseToken(token)

		if claims == nil || err != nil {
			context.JSON(http.StatusInternalServerError, entity.Response{
				Code: 0,
				Data: nil,
				Msg:  nil,
			})
			context.Abort()
			return
		} else {
			// when JwtCheck, if check pass, we can trans EmpId to Context
			context.Set("EmpId", claims.EmpId)
			context.Next()
		}
	}
}
