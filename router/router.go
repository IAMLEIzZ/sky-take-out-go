package router

import (
	"net/http"
	"sky-take-out-go/internal/controller/base"
	"sky-take-out-go/model/entity"
	"sky-take-out-go/internal/controller/admin"
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
		router.POST("/admin/employee", JwtHandler(), admin.SaveEmp)	//  新增员工路由
		router.POST("/admin/employee/login", admin.Login)	//  员工登录路由
		router.GET("/admin/employee/:id", JwtHandler(), admin.GetEmpById)		// 根据ID查询员工路由	
		router.POST("/admin/employee/status/:status", JwtHandler(), admin.SetEmpStatus)		// 启用或停用员工路由
		router.PUT("/admin/employee/editPassword", JwtHandler(), admin.EditPassword)	 // 修改密码路由
		router.POST("/admin/employee/logout", JwtHandler(), admin.EmpLogout)		// 退出登录路由
		router.PUT("/admin/employee", JwtHandler(), admin.EditEmp)		// 修改员工信息路由
		router.GET("/admin/employee/page", JwtHandler(), admin.PageQueryEmp)	//  分页查询员工路由

	}
	
	// 分类管理路由
	{
		router.POST("/admin/category", JwtHandler(), admin.SaveCate)	//  新增分类路由
		router.GET("/admin/category/page", JwtHandler(), admin.PageQueryCate)	//  分页查询分类路由
		router.DELETE("/admin/category", JwtHandler(), admin.DeleteCateById)  // 根据 ID 删除菜品分类
		router.PUT("/admin/category", JwtHandler(), admin.UpdateCate)	// 修改分类路由
		router.POST("/admin/category/status/:status", JwtHandler(), admin.SetCateStatus)	// 启用或停用分类路由
		router.GET("/admin/category/list", JwtHandler(), admin.GetCateListByType)	// 根据 type 查询分类列表
	}
	
	// 菜品管理路由
	{
		router.POST("/admin/dish", JwtHandler(), admin.SaveDish)	//  新增菜品路由
		router.GET("/admin/dish/page", JwtHandler(), admin.DishPageQuery)	//  分页查询菜品路由
		router.DELETE("/admin/dish", JwtHandler(), admin.DeleteDish)  // 批量删除菜品
		router.GET("/admin/dish/:id", JwtHandler(), admin.GetDishById)		// 根据ID查询菜品路由
		router.GET("/admin/dish/list", JwtHandler(), admin.GetDishByCategoryId)	// 根据分类查询菜品列表
		router.PUT("/admin/dish", JwtHandler(), admin.UpdateDish)	// 修改菜品路由
		router.POST("/admin/dish/status/:status", JwtHandler(), admin.SetDishStatus)	// 启用或停用菜品路由
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
