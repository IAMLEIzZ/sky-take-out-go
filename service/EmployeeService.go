package service

import (
	"log"
	"sky-take-out-go/dao"
	"sky-take-out-go/model/dto"
	"sky-take-out-go/model/entity"
	"sky-take-out-go/utils"

	"github.com/jinzhu/copier"
)

func Save(employeeDTO dto.EmployeeDTO) error {

	employee := entity.Employee{}

	err := copier.Copy(&employee, employeeDTO)

	if err != nil {
		log.Println("INFO: Object Copy fail..." + err.Error())
	}

	defaultPassword := "123456"

	employee.Password = utils.Md5DigestAsHex(defaultPassword)

	employee.Status = 1

	return dao.Save(employee)
}

func PageQuery(employeePageQueryDTO dto.EmployeePageQueryDTO) ([]entity.Employee, int64, error) {

	employs, total, err := dao.PageQuery(employeePageQueryDTO)

	return employs, total, err
}
