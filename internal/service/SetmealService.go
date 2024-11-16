package service

import (
	"errors"
	"sky-take-out-go/internal/api/request"
	"sky-take-out-go/internal/api/response"
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

func DeleteSetmealBatch(ids []uint64) error {
	// delete setmeal
	err := dao.DeleteSetmealBatch(ids)
	if err != nil {
		return err
	}
	// delete setmealdish
	err = dao.DeleteSetmealDishBatch(ids)
	return err
}

func UpdateSetmeal(c *gin.Context, setmealDTO *request.SetMealDTO) error {
	price, _ := strconv.ParseFloat(setmealDTO.Price, 64)
	setmeal := &model.SetMeal{
		Id:          setmealDTO.Id,
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
		setmeal.UpdateTime = time.Now()
	} else {
		return errors.New("获取当前用户信息失败")
	}
	// update setmeal
	err := dao.UpdateSetmeal(setmeal)
	if err != nil {
		return err
	}

	// delete setmealdish
	var setmealIds []uint64
	setmealIds = append(setmealIds, setmeal.Id)
	err = dao.DeleteSetmealDishBatch(setmealIds)
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

func GetSetmealById(id uint64) (*response.SetMealWithDishByIdVo , error) {
	setmeal, err := dao.GetSetmealById(id)
	if err != nil {
		return nil, err
	}

	setmealdishlist, err := dao.GetSetmealDishBySetmealId(id)
	if err != nil {
		return nil, err
	}
	setmealVo := &response.SetMealWithDishByIdVo{
		Id: setmeal.Id,
		CategoryId: setmeal.CategoryId,
		Name: setmeal.Name,
		Price: setmeal.Price,
		Status: setmeal.Status,
		Description: setmeal.Description,
		Image: setmeal.Image,
		SetmealDishes: setmealdishlist,
	}
	return setmealVo, nil
}

func SetSetmealStatus(status int, id uint64, c *gin.Context) error {
	setmeal := &model.SetMeal{
		Id: id,
		Status: status,
	}
	// get EmpID
	if empId, exsits := c.Get("EmpId"); exsits {
		setmeal.UpdateUser = empId.(uint64)
		setmeal.UpdateTime = time.Now()
	} else {
		return errors.New("获取当前用户信息失败")
	}
	err := dao.UpdateSetmeal(setmeal)
	if err != nil {
		return err
	}
	err = dao.SetSetmealStatus(id, status)

	return err
}