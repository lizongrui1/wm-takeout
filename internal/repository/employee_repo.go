package repository

import (
	"context"
	"wm-take-out/internal/model"
)

type EmployeeRepo interface {
	GetUserName(ctx context.Context, userName string) (*model.Employee, error)
	GetId(ctx context.Context, id uint64) (*model.Employee, error)
	InsertUser(ctx context.Context, entity model.Employee) error
	UpdateUser(ctx context.Context, employee model.Employee) error
}
