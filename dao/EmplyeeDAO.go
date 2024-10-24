package dao

import (
	"sky-take-out-go/db"
	"sky-take-out-go/model/entity"
)

// 新增一个员工
func Save(employee entity.Employee) error{

	err := db.DB.Create(&employee)

	return err.Error
}
