package entity

import (
    "time"
)

// Category 表示数据库中的分类表
type Category struct {
    ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
    // 类型: 1菜品分类 2套餐分类
    Type       int       `gorm:"type:int;not null" json:"type"`
    // 分类名称
    Name       string    `gorm:"type:varchar(255);not null" json:"name"`
    // 顺序
    Sort       int       `gorm:"type:int;default:0" json:"sort"`
    // 分类状态 0表示禁用 1表示启用
    Status     int       `gorm:"type:int;default:1" json:"status"`
    // 创建时间
    CreateTime time.Time `gorm:"autoCreateTime" json:"create_time"`
    // 更新时间
    UpdateTime time.Time `gorm:"autoUpdateTime" json:"update_time"`
    // 创建人
    CreateUser uint64    `gorm:"type:bigint;not null" json:"create_user"`
    // 修改人
    UpdateUser uint64    `gorm:"type:bigint;not null" json:"update_user"`
}

func (Category) TableName() string {
	return "category" // 指定表名为 category
}