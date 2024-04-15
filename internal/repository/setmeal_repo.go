package repository

import (
	"gorm.io/gorm"
	"wm-take-out/global/tx"
	"wm-take-out/internal/model"
)

type SetMealRepo interface {
	InsertCombo(db tx.Transaction, meal *model.SetMeal) error
	DeleteCombo(db *gorm.DB, id uint64) error
	UpdateCombo
}
