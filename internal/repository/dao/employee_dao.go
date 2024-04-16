package dao

import (
	"context"
	"gorm.io/gorm"
	"wm-take-out/internal/model"
	"wm-take-out/internal/repository"
)

type EmployeeDao struct {
	db *gorm.DB
}

func (e *EmployeeDao) GetByUserName(ctx context.Context, userName string) (*model.Employee, error) {
	var employee model.Employee
	err := e.db.WithContext(ctx).Where("username = ?", userName).First(&employee).Error
	return &employee, err
}

func (e *EmployeeDao) GetById(ctx context.Context, id uint64) (*model.Employee, error) {
	var employee model.Employee
	err := e.db.WithContext(ctx).Where("id = ?", id).First(&employee).Error
	return &employee, err
}

func (e *EmployeeDao) InsertUser(ctx context.Context, user model.Employee) error {
	err := e.db.WithContext(ctx).Create(&user).Error
	return err
}

func (e *EmployeeDao) UpdateUser(ctx context.Context, employee model.Employee) error {
	err := e.db.WithContext(ctx).Model(&employee).Updates(employee).Error
	return err
}

func (e *EmployeeDao) UpdateStatus(ctx context.Context, employee model.Employee) error {
	err := e.db.WithContext(ctx).Model(&employee).Where("id = ?", employee.Id).Update("status", employee.Status).Error
	return err
}

func NewEmployeeDao(db *gorm.DB) repository.EmployeeRepo {
	return &EmployeeDao{db: db}
} //？
