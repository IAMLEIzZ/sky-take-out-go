package service

import (
	"errors"
	"sky-take-out-go/internal/dao"
	"sky-take-out-go/internal/api/request"
	"sky-take-out-go/internal/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveDishWithFlavors(dishdto *request.DishDTO, c *gin.Context) error {
	// trans dto to entity
	tmp_price, err := strconv.ParseFloat(dishdto.Price, 64)
	// err := copier.Copy(dish, dishdto)
	if err != nil {
		return err
	}
	// Step 1: Save Dish
	dish := &model.Dish{
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
	err = dao.InsertDish(dish, c)

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
		err = dao.InsertFlavorBatch(dishFlavors)

		if err != nil {
			return err
		}
	}

	return nil
}

func DishPageQuery(dishPageQueryDTO *request.DishPageQueryDTO) ([]model.Dish, int64, error) {
	dishs, total, err := dao.DishPageQuery(dishPageQueryDTO)
	return dishs, total, err
}

func DeleteDishBatch(ids []uint64) error {
	err := dao.DeleteDishBatch(ids)

	if err != nil {
		return err
	}

	// delete dish flavors by dishid
	for _, id := range ids {
		err = dao.DeleteFlavorByDishId(id)

		if err != nil {
			return err
		}
	}
	
	return err
}

func GetDishById(id uint64) (*model.Dish,  error){
	dish, err := dao.GetDishById(id)
	return dish, err
}

func List(categoryId uint64) ([]model.Dish, error) {
	Dish := &model.Dish{
		CategoryId: categoryId,
		Status: 1,
	}

	dishes, err := dao.DishList(Dish)

	return dishes, err
}

func DishUpdate(dishdto *request.DishDTO, c *gin.Context) error {
	// Update dish && Create dish 
	price, err := strconv.ParseFloat(dishdto.Price, 64)
	if err != nil {
		return err
	}
	dish := &model.Dish{
		Id: dishdto.Id,
		Name: dishdto.Name,
		CategoryId: dishdto.CategoryId,
		Price: price,
		Image: dishdto.Image,
		Description: dishdto.Description,
		Status: dishdto.Status,
	}
	if empId, exists := c.Get("EmpId"); exists {
		dish.UpdateUser = empId.(uint64)
		dish.UpdateTime = time.Now()
	} else {
		return errors.New("admin not exist")
	}
	err = dao.DishUpdate(dish)
	if err != nil {
		return err
	}

	// Update flavors
	// Delete All Flavors By DishID
	err = dao.DeleteFlavorByDishId(dish.Id)
	if err != nil {
		return err
	}

	dishFlavors := dishdto.Flavors
	if(len(dishFlavors) > 0) {
		// Set DishID for Every Flavor
		for i := range dishFlavors {
			dishFlavors[i].DishId = dish.Id
		}
		// save dish
		err = dao.InsertFlavorBatch(dishFlavors)

		if err != nil {
			return err
		}
	}	
	return nil
}

func SetDishStatus(id uint64, status int, c *gin.Context) error {
	dish, err := dao.GetDishById(id)
	if err != nil {
		return err
	}
	dish.Status = status

	if empId, exists := c.Get("EmpId"); exists {
		dish.UpdateUser = empId.(uint64)
		dish.UpdateTime = time.Now()
	} else {
		return errors.New("admin not exist")
	}

	err = dao.UpdateDishSatatus(dish)
	return err
}