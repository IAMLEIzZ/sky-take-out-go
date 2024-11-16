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

func DeleteSetmealDishBatch(ids []uint64) error {
	err := db.DB.Debug().Where("setmeal_id in (?)", ids).Delete(&model.SetMealDish{}).Error
	return err
}