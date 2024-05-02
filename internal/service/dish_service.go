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
	PageQuery(ctx context.Context, dto request.DishPageQueryDTO) (*common.PageResult, error)
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
			transation, _ := ds.repo.Transaction(ctx)
			defer func() {
				if r := recover(); r != nil {
					transation.Rollback()
				}
			}()
			err := ds.repo.Delete(transation, dishId)
			if err != nil {
				return err
			}
			err = ds.dishFlavorRepo.DeleteById(transation, dishId)
			if err != nil {
				return err
			}
		}()
		if err != nil {
			return err
		}
		return nil

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
}

func (ds *DishSe) GetDishById(ctx context.Context, id uint64) (response.DishVo, error) {

}
