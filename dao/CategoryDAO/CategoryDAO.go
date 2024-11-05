package categorydao

import (
	"sky-take-out-go/db"
	"sky-take-out-go/model/dto"
	"sky-take-out-go/model/entity"

	"github.com/spf13/cast"
)

func Save(category *entity.Category) error {
	
	err := db.DB.Debug().Create(category)

	return err.Error 
}

func PageQuery(categoryPageQueryDTO dto.CategoryPageQueryDTO) ([]entity.Category, int64, error) {	
	var categories []entity.Category
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
	query := db.DB.Debug().Model(&entity.Category{})

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