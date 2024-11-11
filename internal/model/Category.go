package model

import (
    "time"
)

// Category 表示数据库中的分类表
type Category struct {
    ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
    Type       int       `gorm:"type:int;not null" json:"type"`
    Name       string    `gorm:"type:varchar(255);not null" json:"name"`
    Sort       int       `gorm:"type:int;default:0" json:"sort"`
    Status     int       `gorm:"type:int;default:1" json:"status"`
    CreateTime time.Time `gorm:"autoCreateTime" json:"create_time"`
    UpdateTime time.Time `gorm:"autoUpdateTime" json:"update_time"`
    CreateUser uint64    `gorm:"type:bigint;not null" json:"create_user"`
    UpdateUser uint64    `gorm:"type:bigint;not null" json:"update_user"`
}

func (Category) TableName() string {
	return "category" // 指定表名为 category
}