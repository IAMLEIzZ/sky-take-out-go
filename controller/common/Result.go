package common

import "sky-take-out-go/model/entity"

type Response struct {
	Code float64     `json:"code"`
	Data interface{}        `json:"data"`
	Msg  *string `json:"msg"`
}

type EmpList struct {
	Total int64 `json:"total"`
	Records []entity.Employee `json:"records"` 
}