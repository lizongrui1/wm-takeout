package service

import (
	"context"
	"wm-take-out/common"
	"wm-take-out/common/e"
	"wm-take-out/common/utils"
	"wm-take-out/internal/api/request"
	"wm-take-out/internal/api/response"
	"wm-take-out/internal/enum"
	"wm-take-out/internal/model"
	"wm-take-out/internal/repository"
)

type EmployeeService interface {
	Login(ctx context.Context, login request.EmployeeLogin) (response.EmployeeLogin, error)
	Logout(ctx context.Context) error
	CreateEmployee(ctx context.Context, dto request.EmployeeDTO) error
	EditPassword(ctx context.Context, word request.EmployeeChangePassWord) error
	EditEmployee(ctx context.Context, dto request.EmployeeDTO) error
	EmployeeQueryById(ctx context.Context, id uint64) (model.Employee, error)
	EmployeeStatus(ctx context.Context, status int) error
	PageQuery(ctx context.Context, dto request.EmployeePageQueryDTO) (*common.PageResult, error)
}

type EmployeeSe struct {
	repo repository.EmployeeRepo
}

func (es *EmployeeSe) Login(ctx context.Context, login request.EmployeeLogin) (*response.EmployeeLogin, error) {
	employee, err := es.repo.GetByUserName(ctx, login.UserName)
	if err != nil || employee == nil {
		return nil, e.Error_ACCOUNT_NOT_FOUND
	}
	password := utils.MD5V(login.Password, "", 0)
	if password != employee.Password {
		return nil, e.Error_PASSWORD_ERROR
	}
	if employee.Status != enum.DISABLE {
		return nil, e.Error_ACCOUNT_LOCKED
	}

}
