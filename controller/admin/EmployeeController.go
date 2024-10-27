package admin

import (
	"log"
	"net/http"
	"sky-take-out-go/controller/common"
	"sky-take-out-go/model/dto"
	"sky-take-out-go/model/vo"
	"sky-take-out-go/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// add a employee
// Path: admin/emplyee
func Save(c *gin.Context) {

	log.Println("INFO: " + "Add a employee")

	employeeDTO := dto.EmployeeDTO{}

	err := c.ShouldBind(&employeeDTO) // 将传入的 JSON 对象赋值给 DTO 对象

	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Error[common.H](err.Error()))
	}

	err = service.Save(employeeDTO)

	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Error[map[string]interface{}](err.Error()))
	}

	c.JSON(http.StatusOK, common.Success[map[string]interface{}]())

}

// page query
// Path: admin/employee/page
func Page(c *gin.Context) {
	log.Println("INFO: " + "Add a employee")
	// 把 context 中的信息绑定到 DTO 中
	employeePageQueryDTO := dto.EmployeePageQueryDTO{}
	err := c.ShouldBind(&employeePageQueryDTO)

	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Error[common.H](err.Error()))
	}

	employees, total, err1 := service.PageQuery(employeePageQueryDTO)

	if err1 != nil {
		c.JSON(http.StatusInternalServerError, common.Error[common.H](err.Error()))
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg": nil,
		"data": gin.H{
			"total": total,
			"records": employees,
		},
    })
}	

type Cliams struct {
	EmpId uint64	`json:"empId"`
	jwt.RegisteredClaims
}


// empolyee login
// Path: /admin/employee/login	
func Login(c *gin.Context) {
	log.Println("INFO: " + "Employee Login")
	employeeLoginDTO := dto.EmployeeLoginDTO{}
	err := c.ShouldBind(&employeeLoginDTO) 	

	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Error[common.H](err.Error()))
		return 
	}

	employee, err := service.Login(employeeLoginDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Error[common.H](err.Error()))
		return 
	}
	// JWT
	JwtAdminSecretKey := "itcast"
	JwtTTL := time.Now().Add(7200000 * time.Second)

	claim := &Cliams{
		EmpId: employee.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(JwtTTL),
		},
	}
	// token 生成	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(JwtAdminSecretKey))	
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Error[common.H](err.Error()))
		return
	}
	employeeLoginVO := vo.EmployeeLoginVO{
		ID: int64(employee.ID),
		UserName: employee.Username,
		Name: employee.Name,
		Token: tokenString,
	}

	c.JSON(http.StatusOK, vo.Response{
		Code: 1,
		Data: &employeeLoginVO,
		Msg: nil,
	})
}



