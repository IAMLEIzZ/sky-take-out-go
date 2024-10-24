package dto


type EmployeePageQueryDTO struct {
	// 员工姓名
	Name string `json:"name" form:"name"`

	// 页码
	Page int `json:"page" form:"page"`

	// 每页显示记录数
	PageSize int `json:"pageSize" form:"pageSize"`
	
}