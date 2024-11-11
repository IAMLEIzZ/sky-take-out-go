package service 

import (
	"errors"
	"log"
	"sky-take-out-go/internal/api/request"
	"sky-take-out-go/internal/dao"
	"sky-take-out-go/internal/model"
	"sky-take-out-go/utils"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func SaveEmp(employeeDTO *request.EmployeeDTO, c *gin.Context) error {

	employee := model.Employee{}

	err := copier.Copy(&employee, employeeDTO)

	if err != nil {
		log.Println("INFO: Object Copy fail..." + err.Error())
	}

	defaultPassword := "123456"

	employee.Password = utils.Md5DigestAsHex(defaultPassword)

	employee.Status = 1
	employee.CreateTime = time.Now()
	employee.UpdateTime = time.Now()
	
	if empId, exsits := c.Get("EmpId"); exsits {
		employee.UpdateUser = empId.(uint64)
		employee.CreateUser = empId.(uint64)
	} else {
		return errors.New("获取当前用户信息失败")
	}

	return dao.SaveEmp(employee)
}

func EmpPageQuery(employeePageQueryDTO request.EmployeePageQueryDTO) ([]model.Employee, int64, error) {

	employs, total, err := dao.EmpPageQuery(employeePageQueryDTO)

	return employs, total, err
}

func Login(employeeLoginDTO request.EmployeeLoginDTO) (model.Employee, error) {
	username := employeeLoginDTO.Username
	password := employeeLoginDTO.Password

	employee := dao.GetEmpByUsername(username)
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

func GetEmpById(EmpId uint64) *model.Employee {
	employee := dao.GetEmpById(EmpId)

	return employee
}

func SetEmpStatus(Status int, EmpId uint64, c *gin.Context) error {
	// equal a update
	employee := dao.GetEmpById(EmpId)
	employee.Status = Status
	employee.UpdateTime = time.Now()

	if empId, exsits := c.Get("EmpId"); exsits {
		employee.UpdateUser = empId.(uint64)
	} else {
		return errors.New("获取用户信息失败")
	}

	err := dao.UpdateEmp(employee)

	return err
}

func EditEmpPassword(empNewAndOldPwDTO *request.EmpNewAndOldPwDTO) error {
	// this obj have oldpw newpw and id
	// check oldpw
	employee := dao.GetEmpById(empNewAndOldPwDTO.EmpId)

	oldPw := utils.Md5DigestAsHex(empNewAndOldPwDTO.OldPassword)
	if oldPw != employee.Password {
		return errors.New("旧密码错误")
	}

	// update this emp
	employee.Password = utils.Md5DigestAsHex(empNewAndOldPwDTO.NewPassword)
	employee.UpdateTime = time.Now()
	err := dao.UpdateEmp(employee)

	return err
}
 
func EditEmp(employeeDTO request.EmployeeDTO, c *gin.Context) error {
	employee := dao.GetEmpById(uint64(employeeDTO.ID))

	if employee.IDNumber == "" {
		return errors.New("员工 ID 有误，请联系管理员")
	}

	err := copier.Copy(&employee, employeeDTO)

	if err != nil {
		log.Println("INFO: Object Copy fail..." + err.Error())
		return err
	}

	if empId, exsits := c.Get("EmpId"); exsits {
		employee.UpdateUser = empId.(uint64)
	} else {
		return errors.New("获取用户信息失败")
	}

	employee.UpdateTime = time.Now()
	return dao.UpdateEmp(employee)
}