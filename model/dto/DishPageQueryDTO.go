package dto

type DishPageQueryDTO struct {
	Page       int    `form:"page"`       // 分页查询的页数
	PageSize   int    `form:"pageSize"`   // 分页查询的页容量
	Name       string `form:"name"`       // 分页查询的name
	CategoryId uint64 `form:"categoryId"` // 分类ID:
	Status     int    `form:"status"`     // 菜品状态
}