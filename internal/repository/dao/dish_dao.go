package dao

import (
	"context"
	"gorm.io/gorm"
	"wm-take-out/internal/model"
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

func (d *DishDao) Commit(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (d *DishDao) RollBack(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (d *DishDao) InsertDish(tx *gorm.DB, dish *model.Dish) error {
	err := tx.Create(dish).Error
	return err
}

func (d *DishDao) Delete(db *gorm.DB, id uint64) error {
	err := db.Delete(&model.Dish{
		Id: id,
	}).Error
	return err
}

func (d *DishDao) Update(db *gorm.DB, dish model.Dish) error {
	err := db.Model(&dish).Updates(dish).Error
	return err
}

func (d *DishDao) GetById(db *gorm.DB, id uint64) (*model.Dish, error) {

}
