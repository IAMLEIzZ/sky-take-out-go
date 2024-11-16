package admin

import (
	"log"
	"sky-take-out-go/internal/api/request"
	"sky-take-out-go/internal/api/response"
	"sky-take-out-go/internal/service"

	"github.com/gin-gonic/gin"
)

// Add a new setmeal
func SaveSetmeal(c *gin.Context) {
	log.Println("INFO: " + "Add a new setmeal")
	setmealDTO := &request.SetMealDTO{}
	err := c.ShouldBindJSON(setmealDTO)
	if err != nil {
		response.Response_Error(c)
		return 
	}

	err = service.AddSetmeal(c, setmealDTO)

	if err != nil {
		response.Response_Error(c)
		return 
	}

	response.Response_Success(c, nil)
}