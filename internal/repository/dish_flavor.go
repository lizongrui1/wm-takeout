package repository

import (
	"gorm.io/gorm"
	"wm-take-out/internal/model"
)

type DishFlavorRepo interface {
	InsertTaste(db *gorm.DB, flavor []model.DishFlavor) error
	DeleteById(db *gorm.DB, id uint64) error
	Update(db *gorm.DB, flavor model.DishFlavor) error
	GetDishFlavor(db *gorm.DB, id uint64) ([]model.DishFlavor, error)
}
