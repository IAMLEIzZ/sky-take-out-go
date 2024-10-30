package dto

type CategoryDTO struct {
    // 分类ID
	ID                   int64 `json:"id,omitempty"`
	// 分类名称                     
	Name                 string `json:"name"`
	// 排序，按照升序排序                
	Sort                 int64  `json:"sort"`
	// 分类类型：1为菜品分类，2为套餐分类       
	Type                 int64  `json:"type"`
}

// CategoryDTOTemp 中间类, 用于接收前端传递的数据，然后将数据转换为CategoryDTO
type CategoryDTOTemp struct {
    // 分类ID  
    ID                   string `json:"id,omitempty"`
    // 分类名称
    Name                 string `json:"name"`
    // 排序，按照升序排序
    Sort                 string `json:"sort"`
    // 分类类型：1为菜品分类，2为套餐分类
    Type                 string `json:"type"`
}