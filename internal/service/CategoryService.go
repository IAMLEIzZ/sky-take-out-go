package service

import (
	"errors"
	"sky-take-out-go/internal/api/request"
	"sky-take-out-go/internal/dao"
	"sky-take-out-go/internal/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)


func SaveCate(categoryDTO *request.CategoryDTO, c *gin.Context) error {

	typ, _ := strconv.Atoi(categoryDTO.Type)
	sort, _ := strconv.Atoi(categoryDTO.Sort)
	category := &model.Category{
		Name: categoryDTO.Name,
		Type: typ,
		Sort: sort,
		Status: 1,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	if empId, exists := c.Get("EmpId"); exists {
		category.CreateUser = empId.(uint64)
		category.UpdateUser = empId.(uint64)
	} else {
		return errors.New("token is invalid")
	}

	err := dao.SaveCate(category)

	return err
}

func CatePageQuery(categoryPageQueryDTO request.CategoryPageQueryDTO) ([]model.Category, int64, error) {
	return dao.CatePageQuery(categoryPageQueryDTO)
}

func GetCateById(Id uint64) *model.Category {
	return dao.GetCateByID(Id)
}

func DeleteCateById(Id uint64) error {
	return dao.DeleteCateById(Id)
}

func UpdateCate(c *gin.Context, categorydto *request.CategoryDTO) error {
	category := dao.GetCateByID(uint64(categorydto.ID))

	if category.Name == "" {
		return errors.New("分类 ID 有误")
	}
	sort, _ := strconv.Atoi(categorydto.Sort)
	category.Sort = sort
	category.Name = categorydto.Name

	if cateId, exists := c.Get("EmpId"); exists {
		category.UpdateUser = cateId.(uint64)
	} else {
		return errors.New("获取用户信息失败") 
	}

	category.UpdateTime = time.Now()
	return dao.CateUpdate(category)
}

func SetCateStatus(status int, Id uint64, c *gin.Context) error {
	category := dao.GetCateByID(Id)
	if category.Name == "" {
		return errors.New("分类 ID 有误")
	}
	category.Status = status
	if cateId, exists := c.Get("EmpId"); exists {
		category.UpdateUser = cateId.(uint64)
	} else {
		return errors.New("获取用户信息失败") 
	}

	category.UpdateTime = time.Now()
	return dao.CateUpdate(category)
}

func CateList(cate_type int64) ([]model.Category, error) {
	return dao.CateList(cate_type)
}