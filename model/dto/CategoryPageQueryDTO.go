package dto

// ShouldBindQuery使用的是结构体的字段名而不是标签中的 json 名称。为了确保所有字段都能正确绑定，可以使用 form 标签来显式映射 URL 中的查询参数

// CategoryPageQueryDTO Go 结构体形式
type CategoryPageQueryDTO struct {
    Page     int    `form:"page"`      // 页码
    PageSize int    `form:"pageSize"`  // 每页记录数
    Name     string `form:"name"`      // 分类名称
    Type     string `form:"type"`      // 分类类型 1 菜品分类 2 套餐分类
}