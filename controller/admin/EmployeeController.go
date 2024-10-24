package admin

import (
	"log"
	"net/http"
	"sky-take-out-go/controller/common"
	"sky-take-out-go/model/dto"
	"sky-take-out-go/service"

	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusInternalServerError, common.Error[common.H](err.Error()))
	}

	c.JSON(http.StatusOK, common.Success[common.H]())

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
		"code": 0,
		"msg": nil,
		"data": gin.H{
			"total": total,
			"records": employees,
		},
    })
}	
