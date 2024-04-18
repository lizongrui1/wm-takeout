package dao

import (
	"context"
	"gorm.io/gorm"
)

type DishDao struct {
	db *gorm.DB
}

func (d *DishDao) Transaction(ctx context.Context) (*gorm.DB, error) {
	tx := d.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}

func (d *DishDao) InsertDish() {
	err := transaction
}
