package common

import (
	"sky-take-out-go/model/entity"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code float64     `json:"code"`
	Data interface{}        `json:"data"`
	Msg  *string `json:"msg"`
}

type EmpList struct {
	Total int64 `json:"total"`
	Records []entity.Employee `json:"records"` 
}

func Response_Error(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, Response{
		Code: 0,
		Msg: nil,
		Data: nil,
	})
}

func Response_Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: 1,
		Data: data,
		Msg: nil,
	})
}