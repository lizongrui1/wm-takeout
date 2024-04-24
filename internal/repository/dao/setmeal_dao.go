package dao

import (
	"context"
	"gorm.io/gorm"
	"wm-take-out/common"
	"wm-take-out/global/tx"
	"wm-take-out/internal/api/request"
	"wm-take-out/internal/model"
	"wm-take-out/internal/repository"
)

type SetMealDao struct {
	db *gorm.DB
}

func (s *SetMealDao) Transaction(ctx context.Context) tx.Transaction {
	return tx.NewGormTransaction(s.db, ctx)
}

func (s *SetMealDao) InsertCombo(trans tx.Transaction, meal *model.SetMeal) error {
	db, err := tx.GetGormDB(trans)
	if err != nil {
		return nil
	}
	err = db.Create(&meal).Error
	return err
}

func (s *SetMealDao) DeleteCombo(ctx context.Context, id uint64) error {
	err := s.db.WithContext(ctx).Model(&model.SetMeal{Id: id}).Update("status", common.StatusInactive).Error
	return err
}

func (s *SetMealDao) PageQuery(ctx context.Context, dto request.SetMealPageQueryDTO) (*common.PageResult, error) {

}
func (s *SetMealDao) SetStatus(ctx context.Context, id uint64, status int) error {
	err := s.db.WithContext(ctx).Model(&model.SetMeal{
		Id: id,
	}).Update("status", status).Error
	return err
}

func NewSetMealDao(db *gorm.DB) repository.SetMealRepo {
	return &SetMealDao{db: db}
}
