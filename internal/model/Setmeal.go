package model

import "time"

type SetMeal struct {
	Id          uint64    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"` // 主键id
	CategoryId  uint64    `json:"category_id"`                         // 分类id
	Name        string    `json:"name"`                                // 套餐名称
	Price       float64   `json:"price"`                               // 套餐单价
	Status      int       `json:"status"`                              // 套餐状态
	Description string    `json:"description"`                         // 套餐描述
	Image       string    `json:"image"`                               // 套餐图片
	CreateTime  time.Time `json:"create_time"`                         // 创建时间
	UpdateTime  time.Time `json:"update_time"`                         // 更新时间
	CreateUser  uint64    `json:"create_user"`                         // 创建用户
	UpdateUser  uint64    `json:"update_user"`                         // 更新用户
}

func (SetMeal) TableName() string {
	return "setmeal" // 指定表名为 setmeal
}