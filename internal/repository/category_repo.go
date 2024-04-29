package repository

import (
	"context"
	"wm-take-out/common"
	"wm-take-out/internal/api/request"
	"wm-take-out/internal/model"
)

type CategoryRepo interface {
	InsertCg(ctx context.Context, category model.Category) error
	DeleteById(ctx context.Context, id uint64) error
	Update(ctx context.Context, category model.Category) error
	PageQuery(ctx context.Context, dto request.CategoryPageQueryDTO) (*common.PageResult, error)
	SetStatus(ctx context.Context, category model.Category) error
	List(ctx context.Context, cate int) ([]model.Category, error)
}
