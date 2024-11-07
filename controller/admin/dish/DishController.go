package dish

import (
	"log"
	"sky-take-out-go/controller/common"
	"sky-take-out-go/model/dto"
	"sky-take-out-go/service/dishservice"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Add a New Dish
// PATH: /admin/dish
func Save(c *gin.Context) {
	log.Println("INFO: " + "Add a New Dish With Flavors")
	// copy dto
	dishDto := &dto.DishDTO{}
	err := c.Bind(dishDto)
	
	if err != nil {
		log.Println("Error : " + err.Error())
		common.Response_Error(c)
		return 
	}

	err = dishservice.SaveWithFlavors(dishDto, c)

	if err != nil {
		log.Println("Error : " + err.Error())	
		common.Response_Error(c)
		return 
	}

	common.Response_Success(c, nil)
}

// Page Query Dish
// PATH: /admin/dish/page
func Page(c *gin.Context) {
	log.Println("INFO: " + "Page Query Dish")
	dishPageQueryDTO := &dto.DishPageQueryDTO{}
	err := c.ShouldBind(dishPageQueryDTO)
	if err != nil {
		common.Response_Error(c)
		return
	}
	dishes, total, err := dishservice.PageQuery(dishPageQueryDTO)
	if err != nil {
		common.Response_Error(c)
		return
	}

	common.Response_Success(c, common.DishList{
		Total: total,
		Records: dishes,
	})
}

// Delete Batch Dish 
// PATH: /admin/dish
func Delete(c *gin.Context) {
	log.Println("INFO: " + "Delete Batch Dish")
	idsParam := c.Query("ids")
	// Parser Param to Array
	idsStr := strings.Split(idsParam, ",")
	var ids []uint64
	for _, idStr := range idsStr {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			common.Response_Error(c)
			return
		}
		ids = append(ids, uint64(id))
	}
	// Delete
	err := dishservice.DeleteBatch(ids)
	if err != nil {
		common.Response_Error(c)
		return
	}
	common.Response_Success(c, nil)
}

// Get Dish By ID
// PATH: /admin/dish
func GetById(c *gin.Context) {
	log.Println("INFO: " + "Get Dish By ID")

	idstr := c.Param("id")
	id, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		common.Response_Error(c)
		return
	}

	dishVo, err := dishservice.GetById(id)
	if err != nil {
		common.Response_Error(c)
		return
	}
	common.Response_Success(c, dishVo)
}