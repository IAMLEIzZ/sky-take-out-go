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
	err := query.Offset((page - 1) * size).Limit(size).Find(&dishes)
	return dishes, total, err.Error
}