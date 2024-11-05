package category

import (
	"log"
	"sky-take-out-go/controller/common"
	"sky-take-out-go/model/dto"
	"strconv"

	"sky-take-out-go/service/categoryservice"

	"github.com/gin-gonic/gin"
)

// Add a category
// PATH: admin/category
func Save(c *gin.Context) {
	log.Println("INFO: " + "Add a category")
	categoryDTOTemp := &dto.CategoryDTOTemp{}
	err1 := c.ShouldBindJSON(&categoryDTOTemp)
	if err1 != nil {
		log.Println("ERROR: " + err1.Error())
		common.Response_Error(c)
		return 
	}
	
	// if err == nil
	sorttmp, err := strconv.ParseInt(categoryDTOTemp.Sort, 10, 64)
	if err != nil {
		log.Println("ERROR: " + err.Error())
		common.Response_Error(c)
		return 
	}
	typeTmp, err := strconv.ParseInt(categoryDTOTemp.Type, 10, 64)
	if err != nil {
		log.Println("ERROR: " + err.Error())
		common.Response_Error(c)
		return 
	}
	categoryDTO := &dto.CategoryDTO{
		Name: categoryDTOTemp.Name,
		Sort: sorttmp,
		Type: typeTmp,
	}
	err = categoryservice.Save(categoryDTO, c)

	if err != nil {
		log.Println("ERROR: " + err.Error())
		common.Response_Error(c)
		return
	}

	common.Response_Success(c, nil)
}


// Page Query Category 
// PATH: admin/category/page
func Page(c *gin.Context) {
	log.Println("INFO: " + "Page Query Category")
	categoryPageQueryDTO := dto.CategoryPageQueryDTO{}
	err := c.ShouldBindQuery(&categoryPageQueryDTO)
	
	if err != nil {
		log.Println("INFO: " + "Json bind error")
		common.Response_Error(c)
		return 
	}

	categorys, totals, err := categoryservice.PageQuery(categoryPageQueryDTO)  

	if err != nil {
		log.Println("ERROR: " + err.Error())
		common.Response_Error(c)
		return 
	}

	common.Response_Success(c, common.CategoryList{
		Total: totals,
		Records: categorys,
	})
}