package categoryservice

import (
	"errors"
	"log"
	"sky-take-out-go/dao/categorydao"
	"sky-take-out-go/model/dto"
	"sky-take-out-go/model/entity"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)


func Save(categoryDTO *dto.CategoryDTO, c *gin.Context) error {

	category := &entity.Category{}
	err := copier.Copy(category, categoryDTO)
	if err != nil {
		return err
	}

	category.Status = 1
	category.CreateTime = time.Now()
	category.UpdateTime = time.Now()
	if empId, exists := c.Get("EmpId"); exists {
		category.CreateUser = empId.(uint64)
		category.UpdateUser = empId.(uint64)
	} else {
		return errors.New("token is invalid")
	}

	err = categorydao.Save(category)

	return err
}

func PageQuery(categoryPageQueryDTO dto.CategoryPageQueryDTO) ([]entity.Category, int64, error) {
	return categorydao.PageQuery(categoryPageQueryDTO)
}

func GetById(Id uint64) *entity.Category {
	return categorydao.GetByID(Id)
}

func DeleteById(Id uint64) error {
	return categorydao.DeleteById(Id)
}

func Update(c *gin.Context, categorydto *dto.CategoryDTO) error {
	category := categorydao.GetByID(uint64(categorydto.ID))

	if category.Name == "" {
		return errors.New("分类 ID 有误")
	}
	type_id := category.Type
	err := copier.Copy(&category, categorydto)
	category.Type = type_id

	if err != nil {
		log.Println("INFO: Object Copy fail..." + err.Error())
		return err
	}

	if cateId, exists := c.Get("EmpId"); exists {
		category.UpdateUser = cateId.(uint64)
	} else {
		return errors.New("获取用户信息失败") 
	}

	category.UpdateTime = time.Now()
	return categorydao.Update(category)
}