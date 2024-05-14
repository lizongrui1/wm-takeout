package dao

import (
	"gorm.io/gorm"
	"wm-take-out/global/tx"
	"wm-take-out/internal/model"
	"wm-take-out/internal/repository"
)

type SetMealDishDao struct {
	db *gorm.DB
}

func (s *SetMealDishDao) InsertCombo(trans tx.Transaction, setmealDishs []model.SetMealDish) error {
	db, err := tx.GetGormDB(trans)
	if err != nil {
		return err
	}
	err = db.Create(setmealDishs).Error
	return err
}

func (s *SetMealDishDao) GetBySetMealId(trans tx.Transaction, Id uint64) ([]model.SetMealDish, error) {
	var SetMealDishList []model.SetMealDish
	db, err := tx.GetGormDB(trans)
	if err != nil {
		return nil, err
	}
	err = db.Where("setmeal_id = ?", Id).Find(&SetMealDishList).Error
	if err != nil {
		return SetMealDishList, err
	}
	return SetMealDishList, err
}

func NewSetmealDishDao() repository.SetMealDishRepo {
	return &SetMealDishDao{}
}
