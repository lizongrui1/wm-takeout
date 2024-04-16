package model

//
//import (
//	"context"
//	"github.com/stretchr/testify/assert"
//	"gorm.io/gorm"
//	"gorm.io/gorm/logger"
//	"modernc.org/sqlite"
//	"testing"
//	"wm-take-out/internal/enum"
//)
//
//func TestCategory_BeforeCreate(t *testing.T) {
//	// 创建一个内存数据库
//	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
//		Logger: logger.Default.LogMode(logger.Silent),
//	})
//	assert.NoError(t, err)
//
//	// 创建一个上下文，并设置当前用户ID
//	ctx := context.WithValue(context.Background(), enum.CurrentId, uint64(1))
//	db = db.WithContext(ctx)
//
//	category := &Category{
//		Name:   "Test Category",
//		Status: 1,
//	}
//
//	// 执行 BeforeCreate
//	err = category.BeforeCreate(db)
//	assert.NoError(t, err)
//
//	// 验证字段是否正确填充
//	assert.NotZero(t, category.CreateTime)
//	assert.NotZero(t, category.UpdateTime)
//	assert.Equal(t, uint64(1), category.CreateUser)
//	assert.Equal(t, uint64(1), category.UpdateUser)
//}
//
//func TestCategory_BeforeUpdate(t *testing.T) {
//	// 创建一个内存数据库
//	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
//		Logger: logger.Default.LogMode(logger.Silent),
//	})
//	assert.NoError(t, err)
//
//	// 创建一个上下文，并设置当前用户ID
//	ctx := context.WithValue(context.Background(), enum.CurrentId, uint64(2))
//	db = db.WithContext(ctx)
//
//	category := &Category{
//		Name:   "Test Category",
//		Status: 1,
//	}
//
//	// 执行 BeforeUpdate
//	err = category.BeforeUpdate(db)
//	assert.NoError(t, err)
//
//	// 验证更新时间和更新用户是否被正确更新
//	assert.NotZero(t, category.UpdateTime)
//	assert.Equal(t, uint64(2), category.UpdateUser)
//}
