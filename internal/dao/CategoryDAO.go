package dao

import (
	"sky-take-out-go/db"
	"sky-take-out-go/internal/api/request"
	"sky-take-out-go/internal/model"
	"github.com/spf13/cast"
)

func SaveCate(category *model.Category) error {
	
	err := db.DB.Debug().Create(category)

	return err.Error 
}

func CatePageQuery(categoryPageQueryDTO request.CategoryPageQueryDTO) ([]model.Category, int64, error) {	
	var categories []model.Category
	var total int64
	// 要支持按照categoryName进行模糊查询，也可以按照categoryType进行查询
	page := cast.ToInt(categoryPageQueryDTO.Page)
	size := cast.ToInt(categoryPageQueryDTO.PageSize)
	// 设置页面
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	// 查询层
	query := db.DB.Debug().Model(&model.Category{})

	// 如果 name 不为空，进行模糊查询
	if name := categoryPageQueryDTO.Name; name != "" {
		query = query.Where("name LIKE ?", "%" + name + "%")
	}

	// 如果 type 不为空，进行匹配
	if cate_type := categoryPageQueryDTO.Type; cate_type != "" {
		query = query.Where("type = ?", cate_type)
	}

	// type and name 都是 nil
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Offset((page - 1) * size).Limit(size).Find(&categories).Error
	return categories, total, err	
}

func GetCateByID(Id uint64) *model.Category {
	category := &model.Category{}

	query := db.DB.Debug().Model(&model.Category{})

	query.Where("id = ?", Id).First(category)

	return category
}

func DeleteCateById(Id uint64) error {

	query := db.DB.Debug().Model(&model.Category{})

	err := query.Delete(&model.Category{}, Id).Error

	return err
}

func CateUpdate(category *model.Category) error {
	query := db.DB.Debug().Model(&model.Category{})

	err := query.Where("id = ?", category.ID).Select("*").Updates(category).Error
	return err
}

func CateList(cate_type int64) ([]model.Category, error) {
	var categories []model.Category

	query := db.DB.Debug().Model(&model.Category{})

	err := query.Where("type = ?", cate_type).Find(&categories).Error

	return categories, err
}