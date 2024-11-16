package dao

import (
	"sky-take-out-go/db"
	"sky-take-out-go/internal/model"
)

func AddSetmeal(setmeal *model.SetMeal) error {
	err := db.DB.Debug().Model(&model.SetMeal{}).Create(setmeal).Error

	return err
}