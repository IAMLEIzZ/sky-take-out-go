package request 

type CategoryDTO struct {
	ID                   int64 `json:"id,omitempty"`
	Name                 string `json:"name"`
	Sort                 string `json:"sort"`
	Type                 string `json:"type"`
}

// CategoryDTOTemp 中间类, 用于接收前端传递的数据，然后将数据转换为CategoryDTO
type CategoryDTOTemp struct {
    ID                   int64 `json:"id,omitempty"`
    Name                 string `json:"name"`
    Sort                 string `json:"sort"`
    Type                 string `json:"type"`
}

type CategoryPageQueryDTO struct {
    Page     int    `form:"page"`      // 页码
    PageSize int    `form:"pageSize"`  // 每页记录数
    Name     string `form:"name"`      // 分类名称
    Type     string `form:"type"`      // 分类类型 1 菜品分类 2 套餐分类
}