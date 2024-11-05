package categoryservice

import (
	"errors"
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