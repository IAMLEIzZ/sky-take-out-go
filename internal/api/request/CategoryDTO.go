package request 

type CategoryDTO struct {
    // 分类ID
	ID                   int64 `json:"id,omitempty"`
	// 分类名称                     
	Name                 string `json:"name"`
	// 排序，按照升序排序                
	Sort                 string `json:"sort"`
	// 分类类型：1为菜品分类，2为套餐分类       
	Type                 string `json:"type"`
}

// CategoryDTOTemp 中间类, 用于接收前端传递的数据，然后将数据转换为CategoryDTO
type CategoryDTOTemp struct {
    // 分类ID  
    ID                   int64 `json:"id,omitempty"`
    // 分类名称
    Name                 string `json:"name"`
    // 排序，按照升序排序
    Sort                 string `json:"sort"`
    // 分类类型：1为菜品分类，2为套餐分类
    Type                 string `json:"type"`
}

type CategoryPageQueryDTO struct {
    Page     int    `form:"page"`      // 页码
    PageSize int    `form:"pageSize"`  // 每页记录数
    Name     string `form:"name"`      // 分类名称
    Type     string `form:"type"`      // 分类类型 1 菜品分类 2 套餐分类
}