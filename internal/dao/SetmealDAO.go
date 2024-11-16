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

func DeleteSetmealBatch(ids []uint64) error {
	err := db.DB.Debug().Where("id in (?)", ids).Delete(&model.SetMeal{}).Error
	return err
}

func UpdateSetmeal(setmeal *model.SetMeal) error {
	err := db.DB.Debug().Model(&model.SetMeal{}).Where("id = ?", setmeal.Id).Updates(setmeal).Error
	return err
}

func GetSetmealById(id uint64) (*model.SetMeal, error) {
	setmeal := &model.SetMeal{}
	err := db.DB.Debug().Model(&model.SetMeal{}).Where("id = ?", id).First(setmeal).Error
	return setmeal, err
}	

func SetSetmealStatus(id uint64, status int) error {
	err := db.DB.Debug().Model(&model.SetMeal{}).Where("id = ?", id).Update("status", status).Error
	return err
}
