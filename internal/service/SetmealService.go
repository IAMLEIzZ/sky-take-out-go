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

func AddSetmeal(c *gin.Context, setmealDTO *request.SetMealDTO) error {
	price, _ := strconv.ParseFloat(setmealDTO.Price, 64)
	setmeal := &model.SetMeal{
		CategoryId:  setmealDTO.CategoryId,
		Name:        setmealDTO.Name,
		Price:       price,
		Status:      setmealDTO.Status,
		Description: setmealDTO.Description,
		Image:       setmealDTO.Image,
	}
	// get EmpID
	if empId, exsits := c.Get("EmpId"); exsits {
		setmeal.UpdateUser = empId.(uint64)
		setmeal.CreateUser = empId.(uint64)
		setmeal.CreateTime = time.Now()
		setmeal.UpdateTime = time.Now()
		setmeal.Status = 1
	} else {
		return errors.New("获取当前用户信息失败")
	}
	// insert setmeal
	err := dao.AddSetmeal(setmeal)
	if err != nil {
		return err
	}

	// set the setmeal ID for each setmealdish
	setmealdishlist := setmealDTO.SetMealDishs
	if(len(setmealdishlist) > 0) {
		for i := range setmealdishlist {
			setmealdishlist[i].SetmealId = setmeal.Id
		}
		// insert setmealdish
		err = dao.InsertBatchSetmealDish(setmealdishlist)

		if err != nil {
			return err
		}
	}

	return nil
}

func SetmealPageQuery(setmealPageQueryDTO request.SetMealPageQueryDTO) ([]model.SetMeal, int64, error) {
	setmeals, total, err := dao.SetmealPageQuery(setmealPageQueryDTO)
	return setmeals, total, err
}