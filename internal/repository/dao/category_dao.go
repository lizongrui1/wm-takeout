package dao

import (
	"context"
	"gorm.io/gorm"
	"wm-take-out/internal/model"
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
