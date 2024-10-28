package dao

import (
	"sky-take-out-go/db"
	"sky-take-out-go/model/dto"
	"sky-take-out-go/model/entity"
	"github.com/spf13/cast"
)

// 新增一个员工
func Save(employee entity.Employee) error{

	err := db.DB.Debug().Create(&employee)

	return err.Error
}

func PageQuery(employeePageQueryDTO dto.EmployeePageQueryDTO) ([]entity.Employee, int64, error) {
	var employees []entity.Employee
	var total int64

	page := cast.ToInt(employeePageQueryDTO.Page)
	size := cast.ToInt(employeePageQueryDTO.PageSize)
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	query := db.DB.Debug().Model(&entity.Employee{})

	// 如果 name 不为空，进行模糊查询
	if name := employeePageQueryDTO.Name; name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	// 统计总条目数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 按照分页要求进行查询
	err := query.Offset((page - 1) * size).Limit(size).Find(&employees).Error
	return employees, total, err
}

// select user by username
func GetByUsername(username string) (entity.Employee) {
	employee := entity.Employee{}
	query := db.DB.Debug().Model(&entity.Employee{})

	query.Where("username = ?", username).First(&employee)

	return employee
}

func GetById(EmpId uint64) *entity.Employee {
	emplyee := &entity.Employee{}

	query := db.DB.Debug().Model(&entity.Employee{})

	query.Where("id = ?", EmpId).First(emplyee)

	return emplyee
}

func Update(employee *entity.Employee) error {
	query := db.DB.Debug().Model(&entity.Employee{})
	// when status = 0
	err := query.Where("id = ?", employee.ID).Select("*").Updates(employee)
	return err.Error
}