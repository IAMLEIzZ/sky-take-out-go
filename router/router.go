package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sky-take-out-go/controller/admin"
	"sky-take-out-go/model/entity"
	"sky-take-out-go/utils"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//  新增员工路由
	router.POST("/admin/employee", JwtHandler(), admin.Save)
	router.GET("/admin/employee/page", JwtHandler(), admin.Page)
	router.POST("/admin/employee/login", admin.Login)
	router.GET("/admin/employee/:id", JwtHandler(), admin.GetById)
	router.POST("/admin/employee/status/:status", JwtHandler(), admin.StartOrStop)
	router.PUT("/admin/employee/editPassword", JwtHandler(), admin.EditPassword)
	router.POST("/admin/employee/logout", JwtHandler(), admin.EmpLogout)

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
