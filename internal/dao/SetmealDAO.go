package dao

import (
	"sky-take-out-go/db"
	"sky-take-out-go/internal/api/request"
	"sky-take-out-go/internal/model"

	"github.com/spf13/cast"
)

func AddSetmeal(setmeal *model.SetMeal) error {
	err := db.DB.Debug().Model(&model.SetMeal{}).Create(setmeal).Error

	return err
}

func SetmealPageQuery(setmealPageQueryDTO request.SetMealPageQueryDTO) ([]model.SetMeal, int64, error) {
	var setmeals []model.SetMeal
	var total int64

	page := cast.ToInt(setmealPageQueryDTO.Page)
	size := cast.ToInt(setmealPageQueryDTO.PageSize)
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	query := db.DB.Debug().Model(&model.SetMeal{})
	if name := setmealPageQueryDTO.Name; name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if categoryId := setmealPageQueryDTO.CategoryId; categoryId != 0 {
		query = query.Where("category_id = ?", categoryId)
	}

	if status := setmealPageQueryDTO.Status; status != 0 {
		query = query.Where("status = ?", status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Offset((page - 1) * size).Limit(size).Find(&setmeals).Error
	return setmeals, total, err
}