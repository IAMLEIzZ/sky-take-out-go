package dto

type CategoryDTO struct {
    ID    int64  `json:"id"`    // 主键
    Type  int    `json:"type"`  // 类型 1 菜品分类 2 套餐分类
    Name  string `json:"name"`  // 分类名称
    Sort  int    `json:"sort"`  // 排序
}