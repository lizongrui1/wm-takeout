package repository

import (
	"wm-take-out/global/tx"
	"wm-take-out/internal/model"
)

type SetMealDishRepo interface {
	InsertCombo(db tx.Transaction, setMealDishs []model.SetMealDish) error
	GetBySetMealId(db tx.Transaction, Id uint64) ([]model.SetMealDish, error)
}
