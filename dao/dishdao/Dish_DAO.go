package dishdao

import (
	"sky-take-out-go/db"
	"sky-take-out-go/model/entity"
	"github.com/gin-gonic/gin"
)

// save a dish
func Insert(dish *entity.Dish, c *gin.Context) error {
	err := db.DB.Debug().Create(dish).Error

	if err != nil {
		return err
	}

	return nil
}