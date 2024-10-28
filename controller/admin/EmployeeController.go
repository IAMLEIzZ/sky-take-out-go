package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"sky-take-out-go/controller/common"
	"sky-take-out-go/model/dto"
	"sky-take-out-go/model/vo"
	"sky-take-out-go/service"
	"sky-take-out-go/utils"
	"strconv"
	"time"
)

// add a employee
// Path: admin/emplyee
func Save(c *gin.Context) {

	log.Println("INFO: " + "Add a employee")

	employeeDTO := dto.EmployeeDTO{}

	err := c.ShouldBind(&employeeDTO) // 将传入的 JSON 对象赋值给 DTO 对象

	if err != nil {
		common.Response_Error(c)
		return
	}

	err = service.Save(employeeDTO)

	if err != nil {
		common.Response_Error(c)
		return
	}

	common.Response_Success(c, nil)
}

// page query
// Path: admin/employee/page
func Page(c *gin.Context) {
	log.Println("INFO: " + "Add a employee")
	// 把 context 中的信息绑定到 DTO 中
	employeePageQueryDTO := dto.EmployeePageQueryDTO{}
	err := c.ShouldBind(&employeePageQueryDTO)

	if err != nil {
		common.Response_Error(c)
		return
	}

	employees, total, err1 := service.PageQuery(employeePageQueryDTO)

	if err1 != nil {
		common.Response_Error(c)
		return
	}

	common.Response_Success(
		c, common.EmpList{
			Total:   total,
			Records: employees,
		})
}

type Cliams struct {
	EmpId uint64 `json:"empId"`
	jwt.RegisteredClaims
}

// empolyee login
// Path: /admin/employee/login
func Login(c *gin.Context) {
	log.Println("INFO: " + "Employee Login")
	employeeLoginDTO := dto.EmployeeLoginDTO{}
	err := c.ShouldBind(&employeeLoginDTO)

	if err != nil {
		common.Response_Error(c)
		return
	}

	employee, err := service.Login(employeeLoginDTO)
	if err != nil {
		common.Response_Error(c)
		return
	}
	// JWT
	JwtTTL := time.Now().Add(7200000 * time.Second)
	claim := &dto.JwtClaimDTO_Admin{
		EmpId: employee.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(JwtTTL),
		},
	}
	jwtAdminSecretKey := dto.JwtAdminSecretKey
	// token 生成
	tokenString, err := utils.CreateJwt(claim, jwtAdminSecretKey)
	if err != nil {
		common.Response_Error(c)
		return
	}
	employeeLoginVO := vo.EmployeeLoginVO{
		ID:       int64(employee.ID),
		UserName: employee.Username,
		Name:     employee.Name,
		Token:    tokenString,
	}

	common.Response_Success(c, &employeeLoginVO)
}

// select user bu uerid
// PATH: /aadmin/employee/:id
func GetById(c *gin.Context) {
	log.Println("INFO: " + "Select User By ID")
	Id := c.Param("id")
	empId, err := strconv.ParseUint(Id, 10, 64)
	if err != nil {
		common.Response_Error(c)
		return
	}
	employee := service.GetById(empId)

	// if employee is nil
	if employee.IDNumber == "" {
		// no user
		common.Response_Error(c)
		return
	}
	// if is not nil
	common.Response_Success(c, employee)

}

// Set Employee Status
// PATH: /admin/employee/status/:status
func StartOrStop(c *gin.Context) {
	log.Println("INFO: " + "Set Employee Status")
	status, err1 := strconv.Atoi(c.Param("status"))
	empId, err2 := strconv.ParseUint(c.Query("id"), 10, 64)
	if err1 != nil || err2 != nil {
		common.Response_Error(c)
		return
	}

	err := service.StartOrStop(status, empId)

	if err != nil {
		common.Response_Error(c)
		return
	}

	common.Response_Success(c, nil)

}

// Edit Password
// PATH: /admin/employee/editPassword
func EditPassword(c *gin.Context) {
	log.Println("INFO: " + "Edit Password")
	var empEditPasswordDTO dto.EmpNewAndOldPwDTO
	// here request just oldPw and newPw
	// so need get EmpID
	err := c.ShouldBindJSON(&empEditPasswordDTO)
	if empId, exsits := c.Get("EmpId"); exsits {
		empEditPasswordDTO.EmpId = empId.(uint64)
	} else {
		common.Response_Error(c)
		return
	}

	if err != nil {
		common.Response_Error(c)
		return
	}

	err = service.EditPassword(&empEditPasswordDTO)

	if err != nil {
		common.Response_Error(c)
		return
	}

	common.Response_Success(c, nil)
}

// User Logout
// PATH: /admin/employee/logout
func EmpLogout(c *gin.Context) {
	log.Println("INFO: " + "User Logout...")

	common.Response_Success(c, nil)

}
