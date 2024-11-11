package middleware

import (
	"sky-take-out-go/internal/api/response"
	"sky-take-out-go/utils"
	"github.com/gin-gonic/gin"
)

func JwtHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("token")
		if token == "" {
			response.Response_Error(context)
			context.Abort()
			return
		}
		claims, err := utils.ParseToken(token)

		if claims == nil || err != nil {
			response.Response_Error(context)
			context.Abort()
			return
		} else {
			// when JwtCheck, if check pass, we can trans EmpId to Context
			context.Set("EmpId", claims.EmpId)
			context.Next()
		}
	}
}
