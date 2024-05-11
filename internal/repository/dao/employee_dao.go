package dao

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"time"
	"wm-take-out/internal/model"
	"wm-take-out/internal/repository"
)

var RedisClient *redis.Client

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

func (e *EmployeeDao) PageQuery(ctx context.Context) error {
	return
}

func NewEmployeeDao(db *gorm.DB) repository.EmployeeRepo {
	return &EmployeeDao{db: db}
} //？

func (e *EmployeeDao) InvalidateToken(ctx context.Context, tokenString string) error {
	// 解析令牌以获取其唯一标识，例如jti claim
	jti, exp, err := parseToken(tokenString)
	if err != nil {
		return err // 或返回自定义错误
	}

	// 计算令牌的剩余有效期
	ttl := time.Unix(exp, 0).Sub(time.Now())

	// 将令牌的jti作为键，存储在Redis中，值可以是任意的，例如"invalidated"
	// 设置与令牌相同的过期时间
	_, err = RedisClient.Set(ctx, jti, "invalidated", ttl).Result()
	if err != nil {
		return err // 或返回自定义错误
	}

	return nil
}

// 解析JWT令牌，获取jti和exp，这里简化处理，你需要用jwt库来实现它
func parseToken(tokenString string) (string, int64, error) {
	// 使用你的JWT库来解析令牌并获取jti和exp
	return "jtiHere", 1234567890, nil
}
