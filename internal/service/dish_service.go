package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"wm-take-out/common"
	"wm-take-out/internal/api/request"
	"wm-take-out/internal/api/response"
	"wm-take-out/internal/enum"
	"wm-take-out/internal/model"
	"wm-take-out/internal/repository"
)

type DishService interface {
	InsertDish(ctx context.Context, dto request.DishDTO) error
	DeleteDish(ctx context.Context, id string) error
	UpdateDish(ctx context.Context, dto request.DishUpdateDTO) error
	GetDishById(ctx context.Context, id uint64) (response.DishVo, error)
	PageQuery(ctx context.Context, dto *request.DishPageQueryDTO) (*common.PageResult, error)
	SetStatus(ctx context.Context, id uint64, status int) error
	List(ctx context.Context, id uint64) ([]response.DishListVo, error)
}

type DishSe struct {
	repo           repository.DishRepo
	dishFlavorRepo repository.DishFlavorRepo
}

func (ds *DishSe) InsertDish(ctx context.Context, dto request.DishDTO) error {
	price, _ := strconv.ParseFloat(strconv.FormatUint(dto.Price, 10), 64) // 注意: 未处理转换错误
	dish := model.Dish{
		Id:          0,
		Name:        dto.Name,
		CategoryId:  dto.CategoryId,
		Price:       price,
		Image:       dto.Image,
		Description: dto.Description,
		Status:      enum.ENABLE,
	}
	transaction, err := ds.repo.Transaction(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	if transaction == nil {
		return fmt.Errorf("transaction is nil without error")
	}
	defer func() {
		if r := recover(); r != nil || transaction.Error != nil {
			transaction.Rollback()
		}
	}()
	if err := ds.repo.InsertDish(transaction, &dish); err != nil {
		return err
	}
	for i := range dto.Flavors {
		dto.Flavors[i].DishId = dish.Id
	}
	if err := ds.dishFlavorRepo.InsertTaste(transaction, dto.Flavors); err != nil {
		return err
	}
	if err := transaction.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (ds *DishSe) DeleteDish(ctx context.Context, id string) error {
	idList := strings.Split(id, ",")
	for _, idList := range idList {
		dishId, _ := strconv.ParseUint(idList, 10, 64)
		err := func() error {
			transaction, _ := ds.repo.Transaction(ctx)
			defer func() {
				if r := recover(); r != nil {
					transaction.Rollback()
				}
			}()
			err := ds.repo.Delete(transaction, dishId)
			if err != nil {
				return err
			}
			err = ds.dishFlavorRepo.DeleteById(transaction, dishId)
			if err != nil {
				return err
			}
			return transaction.Commit().Error
		}()
		if err != nil {
			return err
		}
	}
	return nil
}

func (ds *DishSe) UpdateDish(ctx context.Context, dto request.DishUpdateDTO) error {
	price, _ := strconv.ParseFloat(dto.Price, 64)
	dish := model.Dish{
		Id:          dto.Id,
		Name:        dto.Name,
		CategoryId:  dto.CategoryId,
		Price:       price,
		Image:       dto.Image,
		Description: dto.Description,
		Status:      dto.Status,
		Flavors:     dto.Flavors,
	}
	transaction, _ := ds.repo.Transaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()
	err := ds.repo.Update(transaction, dish)
	if err != nil {
		return err
	}
	err = ds.dishFlavorRepo.DeleteById(transaction, dish.Id)
	if err != nil {
		return err
	}
	if len(dish.Flavors) != 0 {
		err = ds.dishFlavorRepo.InsertTaste(transaction, dish.Flavors)
		if err != nil {
			return err
		}
	}
	return transaction.Commit().Error
}

func (ds *DishSe) GetDishById(ctx context.Context, id uint64) (response.DishVo, error) {
	dish, err := ds.repo.GetById(ctx, id)
	dishVo := response.DishVo{
		Id:          dish.Id,
		Name:        dish.Name,
		CategoryId:  dish.CategoryId,
		Price:       dish.Price,
		Image:       dish.Image,
		Description: dish.Description,
		Status:      dish.Status,
		UpdateTime:  dish.UpdateTime,
		Flavors:     dish.Flavors,
	}
	return dishVo, err
}

func (ds *DishSe) PageQuery(ctx context.Context, dto *request.DishPageQueryDTO) (*common.PageResult, error) {
	pageResult, err := ds.repo.PageQuery(ctx, dto)
	if err != nil {
		return nil, err
	}
	return pageResult, err
}

func (ds *DishSe) SetStatus(ctx context.Context, id uint64, status int) error {
	err := ds.repo.Status(ctx, id, status)
	if err != nil {
		return err
	}
	return nil
}

func (ds *DishSe) List(ctx context.Context, id uint64) ([]response.DishListVo, error) {
	var dishListVo []response.DishListVo
	dishList, err := ds.repo.List(ctx, id)
	if err != nil {
		return nil, err
	}
	for _, dish := range dishList {
		dishListVo = append(dishListVo, response.DishListVo{
			Id:          dish.Id,
			Name:        dish.Name,
			CategoryId:  dish.CategoryId,
			Price:       dish.Price,
			Image:       dish.Image,
			Description: dish.Description,
			Status:      dish.Status,
			CreateTime:  dish.CreateTime,
			UpdateTime:  dish.UpdateTime,
			CreateUser:  dish.CreateUser,
			UpdateUser:  dish.UpdateUser,
		})
	}
	return dishListVo, nil
}

func NewDishSe(repo repository.DishRepo, flavorRepo repository.DishFlavorRepo) DishService {
	return &DishSe{repo: repo, dishFlavorRepo: flavorRepo}
}
