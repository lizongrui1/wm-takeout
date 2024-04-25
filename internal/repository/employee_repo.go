package repository

import (
	"context"
	"wm-take-out/internal/model"
)

type EmployeeRepo interface {
	GetByUserName(ctx context.Context, userName string) (*model.Employee, error)
	GetById(ctx context.Context, id uint64) (*model.Employee, error)
	InsertUser(ctx context.Context, user model.Employee) error
	UpdateUser(ctx context.Context, employee model.Employee) error
	UpdateStatus(ctx context.Context, employee model.Employee) error
	PageQuery(ctx context.Context) error
	InvalidateToken(ctx context.Context, tokenString string) error
}
