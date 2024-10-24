package service

import (
	"github.com/jinzhu/copier"
	"log"
	"sky-take-out-go/dao"
	"sky-take-out-go/model/dto"
	"sky-take-out-go/model/entity"
	"sky-take-out-go/utils"
)

func Save(employeeDTO dto.EmployeeDTO) error {

	employee := entity.Employee{}

	err := copier.Copy(&employee, employeeDTO)

	if err != nil {
		log.Fatal("Object Copy fail...", err)
	}

	defaultPassword := "123456"

	employee.Password = utils.Md5DigestAsHex(defaultPassword)

	employee.Status = 1

	return dao.Save(employee)
}
