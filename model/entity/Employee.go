package entity

import (
	"time"
)

// Employee struct for GORM
type Employee struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Username   string    `gorm:"type:varchar(255);not null" json:"username"`
	Name       string    `gorm:"type:varchar(255);not null" json:"name"`
	Password   string    `gorm:"type:varchar(255);not null" json:"password"`
	Phone      string    `gorm:"type:varchar(20);not null" json:"phone"`
	Sex        string    `gorm:"type:varchar(10);not null" json:"sex"`
	IDNumber   string    `gorm:"type:varchar(20);not null" json:"idNumber"`
	Status     int       `gorm:"type:int;not null" json:"status"`
	CreateTime time.Time `gorm:"autoCreateTime" json:"createTime"` // 自动插入创建时间
	UpdateTime time.Time `gorm:"autoUpdateTime" json:"updateTime"` // 自动更新修改时间
	CreateUser uint64    `gorm:"type:bigint;not null" json:"createUser"`
	UpdateUser uint64    `gorm:"type:bigint;not null" json:"updateUser"`
}

func (Employee) TableName() string {

	return "employee" // 指定表名为 employee
}

type Response struct {
	Code float64 `json:"code"`
	Data *Data   `json:"data"`
	Msg  *string `json:"msg"`
}

type Data struct {
	Records []Employee `json:"records"`
	Total   int64      `json:"total"`
}
