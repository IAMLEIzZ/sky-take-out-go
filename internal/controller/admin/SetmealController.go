package admin

import (
	"log"
	"sky-take-out-go/internal/api/request"
	"sky-take-out-go/internal/api/response"
	"sky-take-out-go/internal/service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Add a new setmeal
// PATH: /admin/setmeal
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


// Page query setmeal
// PATH: /admin/setmeal/page
func PageQuerySetmeal(c *gin.Context) {
	log.Println("INFO: " + "Page query setmeal")
	setmealPageQueryDTO := request.SetMealPageQueryDTO{}
	err := c.ShouldBind(&setmealPageQueryDTO)

	if err != nil {
		response.Response_Error(c)
		return 
	}

	setmeals, total, err := service.SetmealPageQuery(setmealPageQueryDTO)

	if err != nil {
		response.Response_Error(c)
		return 
	}

	response.Response_Success(c, response.SetMealList{
		SetMeals: setmeals,
		Total: total,
	})
}

// Delete setmeal
// PATH: /admin/setmeal
func DeleteSetmeal(c *gin.Context) {
	log.Println("INFO: " + "Delete setmeal")
	idsParam := c.Query("ids")
	// Parser Param to Array
	idsStr := strings.Split(idsParam, ",")
	var ids []uint64
	for _, idStr := range idsStr {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.Response_Error(c)
			return
		}
		ids = append(ids, uint64(id))
	}

	err := service.DeleteSetmealBatch(ids)

	if err != nil {
		response.Response_Error(c)
		return 
	}

	response.Response_Success(c, nil)
}