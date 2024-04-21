package dao

import (
	"gorm.io/gorm"
	"wm-take-out/internal/model"
)

type DishFlavor struct {
	db *gorm.DB
}

func (d *DishFlavor) InsertTaste(db *gorm.DB, flavor []model.DishFlavor) error {
	err := db.Create(&flavor).Error
	return err
}

func (d *DishFlavor) DeleteById(db *gorm.DB, id uint64) error {
	err := db.Where("dish_id = ?", id).Delete(&model.DishFlavor{}).Error
	return err
}

func (d *DishFlavor) Update(db *gorm.DB, flavor model.DishFlavor) error {
	err := db.Model(&model.DishFlavor{
		Id: flavor.Id,
	}).Updates(flavor).Error
	return err
}

func (d *DishFlavor) GetDishFlavor(db *gorm.DB, id uint64) ([]model.DishFlavor, error) {
	var dishFlavors []model.DishFlavor
	err := db.Where("id = ?", id).Find(&dishFlavors).Error
	if err != nil {
		return nil, err
	}
	return dishFlavors, nil
}
