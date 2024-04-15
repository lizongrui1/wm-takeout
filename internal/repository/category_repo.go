package repository

import (
	"context"
	"wm-take-out/internal/model"
)

type CategoryRepo interface {
	InsertCg(ctx context.Context, category model.Category) error
	DeleteById(ctx context.Context, id uint64) error
	Update(ctx context.Context, category model.Category) error
	GetCgById(ctx context.Context, id uint64) (*model.Category, error)
	SetStatus(ctx context.Context, category model.Category) error
}
