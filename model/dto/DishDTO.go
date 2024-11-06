package dto

import "sky-take-out-go/model/entity"

type DishDTO struct {
	Id          uint64             `json:"id"`
	Name        string             `json:"name"`
	CategoryId  uint64             `json:"categoryId"`
	Price       string             `json:"price"`
	Image       string             `json:"image"`
	Description string             `json:"description"`
	Status      int                `json:"status"`
	Flavors     []entity.DishFlavor `json:"flavors"`
}