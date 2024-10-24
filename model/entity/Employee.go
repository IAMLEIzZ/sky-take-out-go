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
	IDNumber   string    `gorm:"type:varchar(20);not null" json:"id_number"`
	Status     int       `gorm:"type:int;not null" json:"status"`
	CreateTime time.Time `gorm:"autoCreateTime" json:"create_time"` // 自动插入创建时间
	UpdateTime time.Time `gorm:"autoUpdateTime" json:"update_time"` // 自动更新修改时间
	CreateUser uint64    `gorm:"type:bigint;not null" json:"create_user"`
	UpdateUser uint64    `gorm:"type:bigint;not null" json:"update_user"`
}

func (Employee) TableName() string {
    return "employee"  // 指定表名为 employee
}