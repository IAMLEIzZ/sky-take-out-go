package request

type EmployeeDTO struct {
    ID       uint64   `gorm:"primaryKey;autoIncrement" json:"id"`        // 映射到数据库的 `id` 字段
    Username string `json:"username"`  // 映射到数据库的 `username` 字段
    Name     string `json:"name"`      // 映射到数据库的 `name` 字段
    Phone    string `json:"phone"`     // 映射到数据库的 `phone` 字段
    Sex      string `json:"sex"`       // 映射到数据库的 `sex` 字段
    IDNumber string `json:"idNumber"`  // 映射到数据库的 `idNumber` 字段
}

// 如果需要表名，可以使用 TableName 方法来指定
func (EmployeeDTO) TableName() string {
    return "employee"  // 指定表名为 employee，如果表名不同则修改
}

type EmpNewAndOldPwDTO struct {
	EmpId uint64	`json:"empId"`
	NewPassword string		`json:"newPassword"` 
	OldPassword string		`json:"oldPassword"`	
}

type EmployeeLoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"` 
}

type EmployeePageQueryDTO struct {
	Name string `json:"name" form:"name"`
	Page int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}
