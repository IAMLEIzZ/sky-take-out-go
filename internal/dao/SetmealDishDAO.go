package dao

import (
	"sky-take-out-go/db"
	"sky-take-out-go/internal/model"
)


func InsertBatchSetmealDish(setmealdishlist []model.SetMealDish) error {
	for i := range setmealdishlist {
		err := db.DB.Debug().Model(&model.SetMealDish{}).Create(&setmealdishlist[i]).Error
		if err != nil {
			return err
		}
	}
	return nil
}