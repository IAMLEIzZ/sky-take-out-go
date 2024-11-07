package router

import (
	"net/http"
	"sky-take-out-go/controller/admin/base"
	"sky-take-out-go/controller/admin/category"
	"sky-take-out-go/controller/admin/dish"
	"sky-take-out-go/controller/admin/employee"
	"sky-take-out-go/model/entity"
	"sky-take-out-go/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// 基础路由
	{
		router.POST("/admin/common/upload", JwtHandler(), base.Upload)	// 上传文件路由
	}
	
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
		router.GET("/admin/category/page", JwtHandler(), category.Page)	//  分页查询分类路由
		router.DELETE("/admin/category", JwtHandler(), category.DeleteById)  // 根据 ID 删除菜品分类
		router.PUT("/admin/category", JwtHandler(), category.Update)	// 修改分类路由
		router.POST("/admin/category/status/:status", JwtHandler(), category.StartOrStop)	// 启用或停用分类路由
		router.GET("/admin/category/list", JwtHandler(), category.GetListByType)	// 根据 type 查询分类列表
	}

	// 菜品管理路由
	{
		router.POST("/admin/dish", JwtHandler(), dish.Save)	//  新增菜品路由
		router.GET("/admin/dish/page", JwtHandler(), dish.Page)	//  分页查询菜品路由
		router.DELETE("/admin/dish", JwtHandler(), dish.Delete)  // 批量删除菜品
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
