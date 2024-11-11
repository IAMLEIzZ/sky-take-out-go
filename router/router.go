package router

import (
	"sky-take-out-go/internal/controller/admin"
	"sky-take-out-go/internal/controller/base"
	"sky-take-out-go/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// 基础路由
	{
		router.POST("/admin/common/upload", middleware.JwtHandler(), base.Upload)	// 上传文件路由
	}
	
	// 员工管理路由
	{
		router.POST("/admin/employee", middleware.JwtHandler(), admin.SaveEmp)	//  新增员工路由
		router.POST("/admin/employee/login", admin.Login)	//  员工登录路由
		router.GET("/admin/employee/:id", middleware.JwtHandler(), admin.GetEmpById)		// 根据ID查询员工路由	
		router.POST("/admin/employee/status/:status", middleware.JwtHandler(), admin.SetEmpStatus)		// 启用或停用员工路由
		router.PUT("/admin/employee/editPassword", middleware.JwtHandler(), admin.EditPassword)	 // 修改密码路由
		router.POST("/admin/employee/logout", middleware.JwtHandler(), admin.EmpLogout)		// 退出登录路由
		router.PUT("/admin/employee", middleware.JwtHandler(), admin.EditEmp)		// 修改员工信息路由
		router.GET("/admin/employee/page", middleware.JwtHandler(), admin.PageQueryEmp)	//  分页查询员工路由

	}
	
	// 分类管理路由
	{
		router.POST("/admin/category", middleware.JwtHandler(), admin.SaveCate)	//  新增分类路由
		router.GET("/admin/category/page", middleware.JwtHandler(), admin.PageQueryCate)	//  分页查询分类路由
		router.DELETE("/admin/category", middleware.JwtHandler(), admin.DeleteCateById)  // 根据 ID 删除菜品分类
		router.PUT("/admin/category", middleware.JwtHandler(), admin.UpdateCate)	// 修改分类路由
		router.POST("/admin/category/status/:status", middleware.JwtHandler(), admin.SetCateStatus)	// 启用或停用分类路由
		router.GET("/admin/category/list", middleware.JwtHandler(), admin.GetCateListByType)	// 根据 type 查询分类列表
	}
	
	// 菜品管理路由
	{
		router.POST("/admin/dish", middleware.JwtHandler(), admin.SaveDish)	//  新增菜品路由
		router.GET("/admin/dish/page", middleware.JwtHandler(), admin.DishPageQuery)	//  分页查询菜品路由
		router.DELETE("/admin/dish", middleware.JwtHandler(), admin.DeleteDish)  // 批量删除菜品
		router.GET("/admin/dish/:id", middleware.JwtHandler(), admin.GetDishById)		// 根据ID查询菜品路由
		router.GET("/admin/dish/list", middleware.JwtHandler(), admin.GetDishByCategoryId)	// 根据分类查询菜品列表
		router.PUT("/admin/dish", middleware.JwtHandler(), admin.UpdateDish)	// 修改菜品路由
		router.POST("/admin/dish/status/:status", middleware.JwtHandler(), admin.SetDishStatus)	// 启用或停用菜品路由
	}
	return router
}
