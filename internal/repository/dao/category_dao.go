package dao

import (
	"context"
	"gorm.io/gorm"
	"wm-take-out/internal/api/request"
	"wm-take-out/internal/model"
	"wm-take-out/internal/repository"
)

type Category struct {
	db *gorm.DB
}

func (c *Category) InsertCg(ctx context.Context, category model.Category) error {
	err := c.db.WithContext(ctx).Create(&category).Error
	return err
}

func (c *Category) DeleteById(ctx context.Context, id uint64) error {
	err := c.db.WithContext(ctx).Delete(id).Error
	return err
}

func (c *Category) Update(ctx context.Context, category model.Category) error {
	err := c.db.WithContext(ctx).Model(&category).Updates(&category).Error
	return err
}

func (c *Category) PageQuery(ctx context.Context, dto request.CategoryPageQueryDTO) error {
	err := c.db.WithContext(ctx).Model()
}

func (c *Category) SetStatus(ctx context.Context, category model.Category) error {
	err := c.db.WithContext(ctx).Model(&category).Update("status", category.Status).Error
	return err
}

func NewCategoryDao(db *gorm.DB) repository.CategoryRepo {
	return &Category{db: db}
}
