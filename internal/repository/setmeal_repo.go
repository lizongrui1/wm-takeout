package repository

import (
	"context"
	"wm-take-out/common"
	"wm-take-out/global/tx"
	"wm-take-out/internal/api/request"
	"wm-take-out/internal/model"
)

type SetMealRepo interface {
	Transaction(ctx context.Context) tx.Transaction
	InsertCombo(tran tx.Transaction, meal *model.SetMeal) error
	DeleteCombo(ctx context.Context, id uint64) error
	PageQuery(ctx context.Context, dto request.SetMealPageQueryDTO) (*common.PageResult, error)
	SetStatus(ctx context.Context, id uint64, status int) error
}
