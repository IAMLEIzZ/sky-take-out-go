package admin

import (
	"log"
	"sky-take-out-go/internal/api/response"
	"sky-take-out-go/internal/api/request"
	"sky-take-out-go/internal/service"
	"strconv"
	"strings"
	"github.com/gin-gonic/gin"
)

// Add a New Dish
// PATH: /admin/dish
func SaveDish(c *gin.Context) {
	log.Println("INFO: " + "Add a New Dish With Flavors")
	// copy dto
	dishDto := &request.DishDTO{}
	err := c.Bind(dishDto)
	
	if err != nil {
		log.Println("Error : " + err.Error())
		response.Response_Error(c)
		return 
	}

	err = service.SaveDishWithFlavors(dishDto, c)

	if err != nil {
		log.Println("Error : " + err.Error())	
		response.Response_Error(c)
		return 
	}

	response.Response_Success(c, nil)
}

// Page Query Dish
// PATH: /admin/dish/page
func DishPageQuery(c *gin.Context) {
	log.Println("INFO: " + "Page Query Dish")
	dishPageQueryDTO := &request.DishPageQueryDTO{}
	err := c.ShouldBind(dishPageQueryDTO)
	if err != nil {
		response.Response_Error(c)
		return
	}
	dishes, total, err := service.DishPageQuery(dishPageQueryDTO)
	if err != nil {
		response.Response_Error(c)
		return
	}

	response.Response_Success(c, response.DishList{
		Total: total,
		Records: dishes,
	})
}

// Delete Batch Dish 
// PATH: /admin/dish
func DeleteDish(c *gin.Context) {
	log.Println("INFO: " + "Delete Batch Dish")
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
	// Delete
	err := service.DeleteDishBatch(ids)
	if err != nil {
		response.Response_Error(c)
		return
	}
	response.Response_Success(c, nil)
}

// Get Dish By ID
// PATH: /admin/dish
func GetDishById(c *gin.Context) {
	log.Println("INFO: " + "Get Dish By ID")

	idstr := c.Param("id")
	id, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		response.Response_Error(c)
		return
	}

	dish, err := service.GetDishById(id)
	dishVo := &request.DishDTO{
		Id: dish.Id,
		Name: dish.Name,
		CategoryId: dish.CategoryId,
		Price: strconv.FormatFloat(dish.Price, 'f', -1, 64),
		Image: dish.Image,
		Description: dish.Description,
		Status: dish.Status,
		Flavors: dish.Flavors,
	}
	if err != nil {
		response.Response_Error(c)
		return
	}
	response.Response_Success(c, dishVo)
}

// Get Dish By CategoryID
// PATH: /admin/dish/list
func GetDishByCategoryId(c *gin.Context) {
	log.Println("INFO: " + "Get Dish By CategoryID")
	categoryIdStr := c.Query("categoryId")
	categoryId, err := strconv.ParseUint(categoryIdStr, 10, 64)
	if err != nil {
		response.Response_Error(c)
		return
	}

	dishes, err := service.List(categoryId)
	if err != nil {
		response.Response_Error(c)
		return
	}	

	response.Response_Success(c, dishes)
}

// Update Dish
// PATH: /admin/dish
func UpdateDish(c *gin.Context) {
	log.Println("INFO: " + "Update Dish")
	dishDto := &request.DishDTO{}
	err := c.ShouldBindJSON(dishDto)
	if err != nil {
		response.Response_Error(c)
		return
	}

	err = service.DishUpdate(dishDto, c)
	if err != nil {
		response.Response_Error(c)
		return
	}

	response.Response_Success(c, nil)
}

// Set Dish Status
// PATH: /admin/dish/status/:status
func SetDishStatus(c *gin.Context) {
	log.Println("INFO: " + "Set Dish Status")
	statusStr := c.Param("status")
	status, err := strconv.Atoi(statusStr)
	if err != nil {
		response.Response_Error(c)
		return
	}
	idStr := c.Query("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	err = service.SetDishStatus(id, status, c)
	if err != nil {
		response.Response_Error(c)
		return
	}

	response.Response_Success(c, nil)
}