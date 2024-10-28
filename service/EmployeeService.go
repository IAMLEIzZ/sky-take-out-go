package service

import (
	"errors"
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

func Login(employeeLoginDTO dto.EmployeeLoginDTO) (entity.Employee, error) {
	username := employeeLoginDTO.Username
	password := employeeLoginDTO.Password

	employee := dao.GetByUsername(username)
	// id nil => employee.IDNumber is ""
	if employee.IDNumber == "" {
		return employee, errors.New("账号不存在")
	}

	// if is not nil, password to md5hax
	password = utils.Md5DigestAsHex(password)

	if password != employee.Password {
		// user password err
		return employee, errors.New("密码错误")
	}

	if employee.Status == 0 {
		return employee, errors.New("账号已被锁定")
	}

	return employee, nil

}

func GetById(EmpId uint64) *entity.Employee {
	employee := dao.GetById(EmpId)

	return employee
}

func StartOrStop(Status int, EmpId uint64) error {
	// equal a update
	employee := dao.GetById(EmpId)
	employee.Status = Status

	err := dao.Update(employee)

	return err
}

func EditPassword(empNewAndOldPwDTO *dto.EmpNewAndOldPwDTO) error {
	// this obj have oldpw newpw and id
	// check oldpw
	employee := dao.GetById(empNewAndOldPwDTO.EmpId)

	oldPw := utils.Md5DigestAsHex(empNewAndOldPwDTO.OldPassword)
	if oldPw != employee.Password {
		return errors.New("旧密码错误")
	}

	// update this emp
	employee.Password = utils.Md5DigestAsHex(empNewAndOldPwDTO.NewPassword)

	err := dao.Update(employee)

	return err
}
