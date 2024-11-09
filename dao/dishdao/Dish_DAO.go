package dishdao

import (
	"sky-take-out-go/db"
	"sky-take-out-go/model/dto"
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

func PageQuery(dishPageQueryDTO *dto.DishPageQueryDTO) ([]entity.Dish, int64, error) {
	page := dishPageQueryDTO.Page
	size := dishPageQueryDTO.PageSize

	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	var dishes []entity.Dish
	var total int64
	query := db.DB.Debug().Model(&entity.Dish{})
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

func DeleteBatch(ids []uint64) error {
	err := db.DB.Debug().Where("id in (?)", ids).Delete(&entity.Dish{}).Error
	return err
}

func GetById(id uint64) (*entity.Dish, error) {
	dish := &entity.Dish{}
	err := db.DB.Debug().Model(&entity.Dish{}).Where("id = ?", id).First(dish).Error
	if err != nil {
		return nil, err
	}
	var dishFlavors []entity.DishFlavor
	// select flavors
	err = db.DB.Debug().Model(&entity.DishFlavor{}).Where("dish_id = ?", id).Find(&dishFlavors).Error
	if err != nil {
		return nil, err 
	}
	dish.Flavors = dishFlavors 

	return dish, nil
}

func List(dish *entity.Dish) ([]entity.Dish, error) {
	dishes := []entity.Dish{}
	err := db.DB.Debug().Model(&entity.Dish{}).Where(dish).Find(&dishes).Error
	return dishes, err
}

func Update(dish *entity.Dish) error {
	err := db.DB.Debug().Model(&entity.Dish{Id: dish.Id}).Updates(dish).Error
	return err
}

func UpdateSatatus(dish *entity.Dish) error {
	err := db.DB.Debug().Model(&entity.Dish{Id: dish.Id}).Update("status", dish.Status).Error
	return err
}