package dto

// CategoryPageQueryDTO Go 结构体形式
type CategoryPageQueryDTO struct {
    Page     int    `json:"page"`      // 页码
    PageSize int    `json:"pageSize"`  // 每页记录数
    Name     string `json:"name"`      // 分类名称
    Type     int    `json:"type"`      // 分类类型 1 菜品分类 2 套餐分类
}