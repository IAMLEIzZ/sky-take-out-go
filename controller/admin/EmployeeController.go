package admin

import (
	"net/http"
	"sky-take-out-go/controller/common"
	"sky-take-out-go/model/dto"
	"sky-take-out-go/service"

	"github.com/gin-gonic/gin"
)

// add a employee
// Path: admin/emplyee
func Save(c *gin.Context) {

	employeeDTO := dto.EmployeeDTO{}

	c.ShouldBind(&employeeDTO) // 将传入的 JSON 对象赋值给 DTO 对象
	err := service.Save(employeeDTO)

	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Error[common.H](err.Error()))
	}

	c.JSON(http.StatusOK, common.Success[common.H]())

}
