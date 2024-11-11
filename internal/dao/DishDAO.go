package dao

import (
	"sky-take-out-go/db"
	"sky-take-out-go/internal/api/request"
	"sky-take-out-go/internal/model"
	"github.com/gin-gonic/gin"
)

// save a dish
func InsertDish(dish *model.Dish, c *gin.Context) error {
	err := db.DB.Debug().Create(dish).Error

	if err != nil {
		return err
	}

	return nil
}

func DishPageQuery(dishPageQueryDTO *request.DishPageQueryDTO) ([]model.Dish, int64, error) {
	page := dishPageQueryDTO.Page
	size := dishPageQueryDTO.PageSize

	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	var dishes []model.Dish
	var total int64
	query := db.DB.Debug().Model(&model.Dish{})
	// if name is not empty, do a fuzzy query
	if name := dishPageQueryDTO.Name; name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	// if categoryId is not empty, do a query
	if categoryId := dishPageQueryDTO.CategoryId; categoryId != 0 {
		query = query.Where("category_id = ?", categoryId)
	}
	// if status is not empty, do a query
	if status := dishPageQueryDTO.Status; status != 0 {
		query = query.Where("status = ?", status)
	}

	// count the total number of items
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// query according to the paging requirements
	err := query.Offset((page - 1) * size).Limit(size).Find(&dishes).Error
	return dishes, total, err
}

func DeleteDishBatch(ids []uint64) error {
	err := db.DB.Debug().Where("id in (?)", ids).Delete(&model.Dish{}).Error
	return err
}

func GetDishById(id uint64) (*model.Dish, error) {
	dish := &model.Dish{}
	err := db.DB.Debug().Model(&model.Dish{}).Where("id = ?", id).First(dish).Error
	if err != nil {
		return nil, err
	}
	var dishFlavors []model.DishFlavor
	// select flavors
	err = db.DB.Debug().Model(&model.DishFlavor{}).Where("dish_id = ?", id).Find(&dishFlavors).Error
	if err != nil {
		return nil, err 
	}
	dish.Flavors = dishFlavors 

	return dish, nil
}

func DishList(dish *model.Dish) ([]model.Dish, error) {
	dishes := []model.Dish{}
	err := db.DB.Debug().Model(&model.Dish{}).Where(dish).Find(&dishes).Error
	return dishes, err
}

func DishUpdate(dish *model.Dish) error {
	err := db.DB.Debug().Model(&model.Dish{Id: dish.Id}).Updates(dish).Error
	return err
}

func UpdateDishSatatus(dish *model.Dish) error {
	err := db.DB.Debug().Model(&model.Dish{Id: dish.Id}).Update("status", dish.Status).Error
	return err
}