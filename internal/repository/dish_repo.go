package repository

import (
	"context"
	"gorm.io/gorm"
	"wm-take-out/common"
	"wm-take-out/internal/api/request"
	"wm-take-out/internal/model"
)

type DishRepo interface {
	Transaction(ctx context.Context) (*gorm.DB, error)
	Commit(tx *gorm.DB) error
	RollBack(tx *gorm.DB) error
	InsertDish(db *gorm.DB, dish *model.Dish) error
	Delete(db *gorm.DB, id uint64) error
	Update(db *gorm.DB, dish model.Dish) error
	GetById(ctx context.Context, id uint64) (*model.Dish, error)
	Status(ctx context.Context, id uint64, status int) error
	PageQuery(ctx context.Context, dto *request.DishPageQueryDTO) (*common.PageResult, error)
	List(ctx context.Context, categoryId uint64) ([]model.Dish, error)
}
