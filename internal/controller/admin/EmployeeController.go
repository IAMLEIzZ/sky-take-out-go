package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"sky-take-out-go/internal/api/request"
	"sky-take-out-go/internal/api/response"
	"sky-take-out-go/internal/service"
	"sky-take-out-go/utils"
	"strconv"
	"time"
)

// add a employee
// Path: admin/emplyee
func SaveEmp(c *gin.Context) {
	log.Println("INFO: " + "Add a employee")

	employeeDTO := request.EmployeeDTO{}

	err := c.ShouldBind(&employeeDTO) // 将传入的 JSON 对象赋值给 DTO 对象

	if err != nil {
		response.Response_Error(c)
		return
	}

	err = service.SaveEmp(&employeeDTO, c)

	if err != nil {
		response.Response_Error(c)
		return
	}

	response.Response_Success(c, nil)
}

// page query
// Path: admin/employee/page
func PageQueryEmp(c *gin.Context) {
	log.Println("INFO: " + "Add a employee")
	// 把 context 中的信息绑定到 DTO 中
	employeePageQueryDTO := request.EmployeePageQueryDTO{}
	err := c.ShouldBind(&employeePageQueryDTO)

	if err != nil {
		response.Response_Error(c)
		return
	}

	employees, total, err1 := service.EmpPageQuery(employeePageQueryDTO)

	if err1 != nil {
		response.Response_Error(c)
		return
	}

	response.Response_Success(
		c, response.EmpList{
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
	employeeLoginDTO := request.EmployeeLoginDTO{}
	err := c.ShouldBind(&employeeLoginDTO)

	if err != nil {
		response.Response_Error(c)
		return
	}

	employee, err := service.Login(employeeLoginDTO)
	if err != nil {
		response.Response_Error(c)
		return
	}
	// JWT
	JwtTTL := time.Now().Add(7200000 * time.Second)
	claim := &request.JwtClaimDTO_Admin{
		EmpId: employee.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(JwtTTL),
		},
	}
	jwtAdminSecretKey := request.JwtAdminSecretKey
	// token 生成
	tokenString, err := utils.CreateJwt(claim, jwtAdminSecretKey)
	if err != nil {
		response.Response_Error(c)
		return
	}
	employeeLoginVO := response.EmployeeLoginVO{
		ID:       int64(employee.ID),
		UserName: employee.Username,
		Name:     employee.Name,
		Token:    tokenString,
	}

	response.Response_Success(c, &employeeLoginVO)
}

// select user bu uerid
// PATH: /aadmin/employee/:id
func GetEmpById(c *gin.Context) {
	log.Println("INFO: " + "Select User By ID")
	Id := c.Param("id")
	empId, err := strconv.ParseUint(Id, 10, 64)
	if err != nil {
		response.Response_Error(c)
		return
	}
	employee := service.GetEmpById(empId)

	// if employee is nil
	if employee.IDNumber == "" {
		// no user
		response.Response_Error(c)
		return
	}
	// if is not nil
	response.Response_Success(c, employee)

}

// Set Employee Status
// PATH: /admin/employee/status/:status
func SetEmpStatus(c *gin.Context) {
	log.Println("INFO: " + "Set Employee Status")
	status, err1 := strconv.Atoi(c.Param("status"))
	empId, err2 := strconv.ParseUint(c.Query("id"), 10, 64)
	if err1 != nil || err2 != nil {
		response.Response_Error(c)
		return
	}

	err := service.SetEmpStatus(status, empId, c)

	if err != nil {
		response.Response_Error(c)
		return
	}

	response.Response_Success(c, nil)

}

// Edit Password
// PATH: /admin/employee/editPassword
func EditPassword(c *gin.Context) {
	log.Println("INFO: " + "Edit Password")
	var empEditPasswordDTO request.EmpNewAndOldPwDTO
	// here request just oldPw and newPw
	// so need get EmpID
	err := c.ShouldBindJSON(&empEditPasswordDTO)
	if empId, exsits := c.Get("EmpId"); exsits {
		empEditPasswordDTO.EmpId = empId.(uint64)
	} else {
		response.Response_Error(c)
		return
	}

	if err != nil {
		response.Response_Error(c)
		return
	}

	err = service.EditEmpPassword(&empEditPasswordDTO)

	if err != nil {
		response.Response_Error(c)
		return
	}

	response.Response_Success(c, nil)
}

// User Logout
// PATH: /admin/employee/logout
func EmpLogout(c *gin.Context) {
	log.Println("INFO: " + "User Logout...")

	response.Response_Success(c, nil)

}

// Edit Employee Info
// PATH: /admin/employee
func EditEmp(c *gin.Context) {
	log.Println("INFO: " + "Edit Employee Info")
	employeeDTO := request.EmployeeDTO{}
	err := c.ShouldBind(&employeeDTO)

	if err != nil {
		response.Response_Error(c)
		return
	}

	err = service.EditEmp(employeeDTO, c)

	if err != nil {
		response.Response_Error(c)
		return
	}

	response.Response_Success(c, nil)
}