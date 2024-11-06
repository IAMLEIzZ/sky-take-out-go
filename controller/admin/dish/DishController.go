package dish

import (
	"log"
	"sky-take-out-go/controller/common"
	"sky-take-out-go/model/dto"
	"sky-take-out-go/service/dishservice"

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