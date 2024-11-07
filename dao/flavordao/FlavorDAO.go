package flavordao

import (
	"sky-take-out-go/db"
	"sky-take-out-go/model/entity"
)

func InsertBatch(flavors []entity.DishFlavor) error {
	for i := range flavors {
		err := db.DB.Debug().Model(&entity.DishFlavor{}).Create(&flavors[i])
		if err.Error != nil {
			return err.Error
		}
	}
	return nil
}