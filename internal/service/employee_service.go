package service

import (
	"context"
	"strings"
	"wm-take-out/common"
	"wm-take-out/common/e"
	"wm-take-out/common/utils"
	"wm-take-out/global"
	"wm-take-out/internal/api/request"
	"wm-take-out/internal/api/response"
	"wm-take-out/internal/enum"
	"wm-take-out/internal/model"
	"wm-take-out/internal/repository"
)

type EmployeeService interface {
	Login(ctx context.Context, login request.EmployeeLogin) (*response.EmployeeLogin, error)
	Logout(ctx context.Context) error
	AddEmployee(ctx context.Context, dto request.EmployeeDTO) error
	UpdatePassword(ctx context.Context, word request.EmployeeChangePassword) error
	UpdateEmployee(ctx context.Context, dto request.EmployeeDTO) error
	EmployeeQueryById(ctx context.Context, id uint64) (model.Employee, error)
	EmployeeStatus(ctx context.Context, id uint64, status int) error
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

	// 生成Token
	jwtConfig := global.Config.Jwt.Admin
	token, err := utils.GenerateToken(employee.Id, jwtConfig.Name, jwtConfig.Secret)
	if err != nil {
		return nil, err
	}
	// 构造返回数据
	resp := response.EmployeeLogin{
		Id:       employee.Id,
		Name:     employee.Name,
		Token:    token,
		UserName: employee.UserName,
	}
	return &resp, nil
}

func (es *EmployeeSe) Logout(ctx context.Context) error {
	// 假设JWT令牌在请求的Authorization头部中发送，并且为"Bearer {token}"
	authHeader := ctx.Value("Authorization").(string)
	if authHeader == "" {
		return e.Error_NO_AUTH_HEADER
	}
	// 提取令牌部分，通常是"Bearer {token}"格式
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	// 假设有一个黑名单或令牌存储机制来标记失效的令牌
	// 这里只是示例，实际应用中可能要使用`es.repo`调用具体的失效逻辑
	err := es.repo.InvalidateToken(ctx, tokenString)
	if err != nil {
		return e.Error_FAILED_TO_LOGOUT
	}
	return nil
}

func (es *EmployeeSe) AddEmployee(ctx context.Context, dto request.EmployeeDTO) error {
	user := model.Employee{
		Id:       dto.Id,
		IdNumber: dto.IdNumber,
		Name:     dto.Name,
		Phone:    dto.Phone,
		Sex:      dto.Sex,
		UserName: dto.UserName,
	}
	user.Status = enum.ENABLE
	// 新增用户初始密码为123456
	user.Password = utils.MD5V("123456", "", 0)
	err := es.repo.InsertUser(ctx, user)
	return err
}

func (es *EmployeeSe) UpdatePassword(ctx context.Context, word request.EmployeeChangePassword) error {
	employee, err := es.repo.GetById(ctx, word.EmpId)
	if err != nil {
		return err
	}
	if employee == nil {
		return e.Error_ACCOUNT_NOT_FOUND
	}
	oldHashPassword := utils.MD5V(word.OldPassword, "", 0)
	if employee.Password != oldHashPassword {
		return e.Error_PASSWORD_ERROR
	}
	// 修改员工密码
	newHashPassword := utils.MD5V(word.NewPassword, "", 0) // 使用新密码生成哈希值
	err = es.repo.UpdateUser(ctx, model.Employee{
		Id:       word.EmpId,
		Password: newHashPassword,
	})
	return err
}

func (es *EmployeeSe) UpdateEmployee(ctx context.Context, dto request.EmployeeDTO) error {
	err := es.repo.UpdateUser(ctx, model.Employee{
		Id:       dto.Id,
		IdNumber: dto.IdNumber,
		Name:     dto.Name,
		Phone:    dto.Phone,
		Sex:      dto.Sex,
		UserName: dto.Name,
	})
	return err
}

func (es *EmployeeSe) EmployeeQueryById(ctx context.Context, id uint64) (model.Employee, error) {
	employee, err := es.repo.GetById(ctx, id)
	//employee.Password = "***"  // 在真实密码的上下文中屏蔽密码
	return *employee, err
}

func (es *EmployeeSe) EmployeeStatus(ctx context.Context, id uint64, status int) error {
	err := es.repo.UpdateStatus(ctx, model.Employee{
		Id:     id,
		Status: status,
	})
	return err
}

func (es *EmployeeSe) PageQuery(ctx context.Context, dto request.EmployeePageQueryDTO) (*common.PageResult, error) {

}

func NewEmployeeService(repo repository.EmployeeRepo) EmployeeService {
	return &EmployeeSe{repo: repo}
}
