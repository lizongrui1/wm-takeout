package dao

import (
	"context"
	"google.golang.org/genproto/googleapis/cloud/common"
	"gorm.io/gorm"
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

func (s *SetMealDao) InsertCombo(db tx.Transaction, meal *model.SetMeal) error {

}

func (s *SetMealDao) DeleteCombo(db *gorm.DB, id uint64) error {

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
