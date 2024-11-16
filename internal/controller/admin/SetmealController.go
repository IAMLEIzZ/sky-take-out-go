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

// Update setmeal
// PATH: /admin/setmeal
func UpdateSetmeal(c *gin.Context) {
	log.Println("INFO: " + "Update setmeal")
	setmealDTO := &request.SetMealDTO{}
	err := c.ShouldBindJSON(setmealDTO)
	if err != nil {
		response.Response_Error(c)
		return 
	}

	err = service.UpdateSetmeal(c, setmealDTO)

	if err != nil {
		response.Response_Error(c)
		return 
	}

	response.Response_Success(c, nil)
}

// Get setmeal by ID
// PATH: /admin/setmeal/:id
func GetSetmealById(c *gin.Context) {
	log.Println("INFO: " + "Get setmeal by ID")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Response_Error(c)
		return 
	}

	setmeal, err := service.GetSetmealById(uint64(id))

	if err != nil {
		response.Response_Error(c)
		return 
	}

	response.Response_Success(c, setmeal)
}

// Set setmeal status
// PATH: /admin/setmeal/status/:status
func SetSetmealStatus(c *gin.Context) {
	log.Println("INFO: " + "Set setmeal status")
	statusStr := c.Param("status")
	status, err := strconv.Atoi(statusStr)
	if err != nil {
		response.Response_Error(c)
		return 
	}
	idStr := c.Query("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	err = service.SetSetmealStatus(status, id, c)
	if err != nil {
		response.Response_Error(c)
		return 
	}

	response.Response_Success(c, nil)
}