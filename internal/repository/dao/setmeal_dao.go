package dao

import (
	"context"
	"gorm.io/gorm"
	"wm-take-out/global/tx"
	"wm-take-out/internal/repository"
)

type SetMealDao struct {
	db *gorm.DB
}

func (s *SetMealDao) Transaction(ctx context.Context) tx.Transaction {

}

func NewSetMealDao(db *gorm.DB) repository.SetMealRepo {
	return &SetMealDao{db: db}
}
