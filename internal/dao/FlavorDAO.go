package dao

import (
	"sky-take-out-go/db"
	"sky-take-out-go/internal/model"
)

func InsertFlavorBatch(flavors []model.DishFlavor) error {
	for i := range flavors {
		err := db.DB.Debug().Model(&model.DishFlavor{}).Create(&flavors[i])
		if err.Error != nil {
			return err.Error
		}
	}
	return nil
}

func DeleteFlavorByDishId(dishId uint64) error {
	err := db.DB.Debug().Where("dish_id = ?", dishId).Delete(&model.DishFlavor{}).Error

	return err 
}