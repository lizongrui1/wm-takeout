package service

import (
	"context"
	"wm-take-out/common"
	"wm-take-out/internal/api/request"
	"wm-take-out/internal/api/response"
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
