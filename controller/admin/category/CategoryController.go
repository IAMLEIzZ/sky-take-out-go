package category

import (
	"log"
	"sky-take-out-go/controller/common"
	"sky-take-out-go/model/dto"
	"sky-take-out-go/service/categoryservice"
	"strconv"

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

// Delete By CateId
// PATH: admin/category
func DeleteById(c *gin.Context) {
	log.Println("INFO: " + "Delete Category By Id")
	Id := c.Query("id")

	categoryId, err := strconv.ParseUint(Id, 10, 64)

	if err != nil {
		log.Println("ERROR: " + err.Error())
		common.Response_Error(c)
		return 
	}

	err = categoryservice.DeleteById(categoryId)

	if err != nil {
		common.Response_Error(c)
		return 
	}

	log.Println("INFO: Successfully deleted category")
	common.Response_Success(c, nil)
}

// UpDate Category
// PATH: /admin/category
func Update(c *gin.Context) {
	log.Println("INFO: " + "Update Category")
	categoryDtoTemp:= &dto.CategoryDTOTemp{}
	err := c.ShouldBind(categoryDtoTemp)	// bind temp json
	if err != nil {
		log.Println("ERROR: " + err.Error())
		common.Response_Error(c)
		return 
	}
	// trans dtotemp to dto
	sortTmp, err := strconv.ParseInt(categoryDtoTemp.Sort, 10, 64)
	if err != nil {
		log.Println("ERROR: " + err.Error())
		common.Response_Error(c)
		return
	}
	// create new dto
	categoryDto := &dto.CategoryDTO{
		ID: categoryDtoTemp.ID,
		Name: categoryDtoTemp.Name,
		Sort: sortTmp,
	}
	err = categoryservice.Update(c, categoryDto)

	if err != nil {
		log.Println("ERROR: " + err.Error())
		common.Response_Error(c)
		return 
	}

	common.Response_Success(c, nil)
}

func StartOrStop(c *gin.Context) {
	log.Println("INFO: " + "Set Category Status")
	status, err := strconv.Atoi(c.Param("status"))
	if err != nil {
		log.Println("ERROR: " + err.Error())
		common.Response_Error(c)
		return
	}
	cateId, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		log.Println("ERROR: " + err.Error())
		common.Response_Error(c)
		return
	}
	err = categoryservice.StartOrStop(status, cateId, c)

	if err != nil {
		log.Println("ERROR: " + err.Error())
		common.Response_Error(c)
		return
	}

	common.Response_Success(c, nil)
}

// Get Category By Id
// PATH: /admin/category
func GetById(c *gin.Context) {
	log.Println("INFO: " + "Get Category By Id")
	Id := c.Query("id")
	categoryId, err := strconv.ParseUint(Id, 10, 64)
	if err != nil {
		log.Println("ERROR: " + err.Error())
		common.Response_Error(c)
		return
	}
	category := categoryservice.GetById(categoryId)

	if category.Name == "" {
		common.Response_Error(c)
		return
	}

	common.Response_Success(c, category)
}