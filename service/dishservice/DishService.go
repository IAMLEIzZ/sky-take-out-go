package dishservice

import (
	"errors"
	"sky-take-out-go/dao/dishdao"
	"sky-take-out-go/dao/flavordao"
	"sky-take-out-go/model/dto"
	"sky-take-out-go/model/entity"
	"sky-take-out-go/model/vo"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveWithFlavors(dishdto *dto.DishDTO, c *gin.Context) error {
	// trans dto to entity
	tmp_price, err := strconv.ParseFloat(dishdto.Price, 64)
	// err := copier.Copy(dish, dishdto)
	if err != nil {
		return err
	}
	// Step 1: Save Dish
	dish := &entity.Dish{
		Name: dishdto.Name,
		CategoryId: dishdto.CategoryId,
		Price: tmp_price,
		Image: dishdto.Image,
		Description: dishdto.Description,
		Status: 1,
	}
	if empId, exists := c.Get("EmpId"); exists {
		dish.CreateUser = empId.(uint64)
		dish.UpdateUser = empId.(uint64)
		dish.CreateTime = time.Now()
		dish.UpdateTime = time.Now()
	} else {
		return errors.New("admin not exist")
	}
	err = dishdao.Insert(dish, c)

	if err != nil {
		return err
	}

	// Step 2: Save Flavors
	// Set DishID for Every Flavor
	dishFlavors := dishdto.Flavors
	if(len(dishFlavors) > 0) {
		for i := range dishFlavors {
			dishFlavors[i].DishId = dish.Id
		}
		// save dish
		err = flavordao.InsertBatch(dishFlavors)

		if err != nil {
			return err
		}
	}

	return nil
}

func PageQuery(dishPageQueryDTO *dto.DishPageQueryDTO) ([]entity.Dish, int64, error) {
	dishs, total, err := dishdao.PageQuery(dishPageQueryDTO)
	return dishs, total, err
}

func DeleteBatch(ids []uint64) error {
	err := dishdao.DeleteBatch(ids)
	return err
}

func GetById(id uint64) (*vo.DishVo,  error){
	dishVo, err := dishdao.GetById(id)
	return dishVo, err
}
