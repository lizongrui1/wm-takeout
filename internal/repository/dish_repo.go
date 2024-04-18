package repository

import (
	"context"
	"gorm.io/gorm"
	"wm-take-out/internal/model"
)

type DishRepo interface {
	Transaction(ctx context.Context) (*gorm.DB, error)
	InsertDish(db *gorm.DB, dish *model.Dish) error
	Delete(db *gorm.DB, id uint64) error
	Update(db *gorm.DB, dish model.Dish) error
	GetById(ctx context.Context, id uint64) (*model.Dish, error)
	Status(ctx context.Context, id uint64, status model.Dish) error
}
