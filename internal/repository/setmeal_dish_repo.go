package repository

import (
	"wm-take-out/global/tx"
	"wm-take-out/internal/model"
)

type SetMealDishRepo interface {
	InsertCombo(trans tx.Transaction, setmealDishs []model.SetMealDish) error
	GetBySetMealId(trans tx.Transaction, Id uint64) ([]model.SetMealDish, error)
}
