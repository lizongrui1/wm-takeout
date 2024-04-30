package service

import (
	"context"
	"strconv"
	"wm-take-out/common"
	"wm-take-out/internal/api/request"
	"wm-take-out/internal/api/response"
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
	price, _ := strconv.ParseFloat(dto.Price, 10)
	dish := model.Dish{
		Id:   0,
		Name: dto.Name,
	}
}

func (ds *DishSe) DeleteDish(ctx context.Context, id string) error {

}

func (ds *DishSe) UpdateDish(ctx context.Context, dto request.DishUpdateDTO) error {

}

func (ds *DishSe) GetDishById(ctx context.Context, id uint64) (response.DishVo, error) {

}
