package service

import (
	"context"
	"fmt"
	"strconv"
	"wm-take-out/common"
	"wm-take-out/internal/api/request"
	"wm-take-out/internal/api/response"
	"wm-take-out/internal/enum"
	"wm-take-out/internal/model"
	"wm-take-out/internal/repository"
)

type SetMealService interface {
	EditSetMeal(ctx context.Context, dto request.SetMealDTO) error
	PageQuery(ctx context.Context, dto request.SetMealPageQueryDTO) (*common.PageResult, error)
	SetStatus(ctx context.Context, id uint64, status int) error
	GetById(ctx context.Context, id uint64) (response.SetMealWithDishByIdVo, error)
}

type SetMealSe struct {
	repo     repository.SetMealRepo
	dishrepo repository.SetMealDishRepo
}

func (ss *SetMealSe) EditSetMeal(ctx context.Context, dto request.SetMealDTO) error {
	price, err := strconv.ParseFloat(dto.Price, 64)
	if err != nil {
		return fmt.Errorf("无法解析价格 '%s': %v", dto.Price, err)
	}
	setmeal := model.SetMeal{
		Id:          dto.Id,
		CategoryId:  dto.CategoryId,
		Name:        dto.Name,
		Price:       price,
		Status:      enum.ENABLE,
		Description: dto.Description,
		Image:       dto.Image,
	}
	transaction := ss.repo.Transaction(ctx)
	err = transaction.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err := recover(); err != nil {
			transaction.Rollback()
		}
	}()
	err = ss.repo.InsertSetMeal(transaction, &setmeal)
	if err != nil {
		return err
	}
	for _, setmealDish := range dto.SetMealDishs {
		setmealDish.SetmealId = setmeal.Id
	}
	err = ss.dishrepo.InsertCombo(transaction, dto.SetMealDishs)
	if err != nil {
		return err
	}
	return transaction.Commit()
}

func (ss *SetMealSe) PageQuery(ctx context.Context, dto request.SetMealPageQueryDTO) (*common.PageResult, error) {
	result, err := ss.repo.PageQuery(ctx, dto)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ss *SetMealSe) SetStatus(ctx context.Context, id uint64, status int) error {
	err := ss.repo.SetStatus(ctx, id, status)
	return err
}

func (ss *SetMealSe) GetById(ctx context.Context, id uint64) (response.SetMealWithDishByIdVo, error) {
	transaction := ss.repo.Transaction(ctx)
	err := transaction.Begin()
	if err != nil {
		return response.SetMealWithDishByIdVo{}, err
	}
	defer func() {
		if err := recover(); err != nil {
			transaction.Rollback()
		}
	}()
	setMeal, err := ss.repo.GetById(transaction, id)
	if err != nil {
		return response.SetMealWithDishByIdVo{}, err
	}
	dishList, err := ss.dishrepo.GetBySetMealId(transaction, id)
	if err != nil {
		return response.SetMealWithDishByIdVo{}, err
	}
	//TODO  检查一下两个Name
	setMealVo := response.SetMealWithDishByIdVo{
		Id:            setMeal.Id,
		CategoryId:    setMeal.CategoryId,
		CategoryName:  setMeal.Name, //？
		Description:   setMeal.Description,
		Image:         setMeal.Image,
		Name:          setMeal.Name, //？
		Price:         setMeal.Price,
		SetmealDishes: dishList,
		Status:        setMeal.Status,
		UpdateTime:    setMeal.UpdateTime,
	}
	return setMealVo, nil
}

func NewSetMealService(repo repository.SetMealRepo, dishrepo repository.SetMealDishRepo) SetMealService {
	return &SetMealSe{repo: repo, dishrepo: dishrepo}
}