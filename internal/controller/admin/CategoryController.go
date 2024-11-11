package admin

import (
	"log"
	"sky-take-out-go/internal/api/request"
	"sky-take-out-go/internal/api/response"
	"sky-take-out-go/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Add a category
// PATH: admin/category
func SaveCate(c *gin.Context) {
	log.Println("INFO: " + "Add a category")
	categoryDTOTemp := &request.CategoryDTOTemp{}
	err := c.ShouldBindJSON(&categoryDTOTemp)
	if err != nil {
		log.Println("ERROR: " + err.Error())
		response.Response_Error(c)
		return
	}

	// if err == nil
	// sorttmp, err := strconv.ParseInt(categoryDTOTemp.Sort, 10, 64)
	// if err != nil {
	// 	log.Println("ERROR: " + err.Error())
	// 	common.Response_Error(c)
	// 	return
	// }
	// typeTmp, err := strconv.ParseInt(categoryDTOTemp.Type, 10, 64)
	// if err != nil {
	// 	log.Println("ERROR: " + err.Error())
	// 	common.Response_Error(c)
	// 	return
	// }
	categoryDTO := &request.CategoryDTO{
		Name: categoryDTOTemp.Name,
		Sort: categoryDTOTemp.Sort,
		Type: categoryDTOTemp.Type,
	}
	err = service.SaveCate(categoryDTO, c)

	if err != nil {
		log.Println("ERROR: " + err.Error())
		response.Response_Error(c)
		return
	}

	response.Response_Success(c, nil)
}

// Page Query Category
// PATH: admin/category/page
func PageQueryCate(c *gin.Context) {
	log.Println("INFO: " + "Page Query Category")
	categoryPageQueryDTO := request.CategoryPageQueryDTO{}
	err := c.ShouldBindQuery(&categoryPageQueryDTO)

	if err != nil {
		log.Println("INFO: " + "Json bind error")
		response.Response_Error(c)
		return
	}

	categorys, totals, err := service.CatePageQuery(categoryPageQueryDTO)

	if err != nil {
		log.Println("ERROR: " + err.Error())
		response.Response_Error(c)
		return
	}

	response.Response_Success(c, response.CategoryList{
		Total:   totals,
		Records: categorys,
	})
}

// Delete By CateId
// PATH: admin/category
func DeleteCateById(c *gin.Context) {
	log.Println("INFO: " + "Delete Category By Id")
	Id := c.Query("id")

	categoryId, err := strconv.ParseUint(Id, 10, 64)

	if err != nil {
		log.Println("ERROR: " + err.Error())
		response.Response_Error(c)
		return
	}

	err = service.DeleteCateById(categoryId)

	if err != nil {
		response.Response_Error(c)
		return
	}

	log.Println("INFO: Successfully deleted category")
	response.Response_Success(c, nil)
}

// UpDate Category
// PATH: /admin/category
func UpdateCate(c *gin.Context) {
	log.Println("INFO: " + "Update Category")
	categoryDTO := &request.CategoryDTO{}
	err := c.ShouldBind(categoryDTO) // bind temp json
	if err != nil {
		log.Println("ERROR: " + err.Error())
		response.Response_Error(c)
		return
	}
	err = service.UpdateCate(c, categoryDTO)

	if err != nil {
		log.Println("ERROR: " + err.Error())
		response.Response_Error(c)
		return
	}

	response.Response_Success(c, nil)
}

func SetCateStatus(c *gin.Context) {
	log.Println("INFO: " + "Set Category Status")
	status, err := strconv.Atoi(c.Param("status"))
	if err != nil {
		log.Println("ERROR: " + err.Error())
		response.Response_Error(c)
		return
	}
	cateId, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		log.Println("ERROR: " + err.Error())
		response.Response_Error(c)
		return
	}
	err = service.SetCateStatus(status, cateId, c)

	if err != nil {
		log.Println("ERROR: " + err.Error())
		response.Response_Error(c)
		return
	}

	response.Response_Success(c, nil)
}

// Get Category By Id
// PATH: /admin/category
func GetCateById(c *gin.Context) {
	log.Println("INFO: " + "Get Category By Id")
	Id := c.Query("id")
	categoryId, err := strconv.ParseUint(Id, 10, 64)
	if err != nil {
		log.Println("ERROR: " + err.Error())
		response.Response_Error(c)
		return
	}
	category := service.GetCateById(categoryId)

	if category.Name == "" {
		response.Response_Error(c)
		return
	}

	response.Response_Success(c, category)
}

// Get Category List By Type
// PATH: /admin/category/list
func GetCateListByType(c *gin.Context) {
	log.Println("INFO: " + "Get Category List By Type")
	categoryType := c.Query("type")
	typeTmp, err := strconv.ParseInt(categoryType, 10, 64)
	if err != nil {
		log.Println("ERROR: " + err.Error())
		response.Response_Error(c)
		return
	}
	categories, err := service.CateList(typeTmp)

	if err != nil {
		log.Println("ERROR: " + err.Error())
		response.Response_Error(c)
		return
	}

	response.Response_Success(c, categories)
}
