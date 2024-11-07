package vo

import (
	"sky-take-out-go/model/entity"
	"time"
)

type DishVo struct {
	Id          uint64    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Name        string    `json:"name"`
	CategoryId  uint64    `json:"categoryId"`
	Price       float64   `json:"price"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	UpdateTime  time.Time `json:"updateTime"`
	Flavors []entity.DishFlavor `json:"flavors"`
}