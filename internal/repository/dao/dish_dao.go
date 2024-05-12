package dao

import (
	"context"
	"gorm.io/gorm"
	"wm-take-out/common"
	"wm-take-out/internal/api/request"
	"wm-take-out/internal/api/response"
	"wm-take-out/internal/model"
	"wm-take-out/internal/repository"
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

func (d *DishDao) GetById(ctx context.Context, id uint64) (*model.Dish, error) {
	var dish model.Dish
	err := d.db.WithContext(ctx).Preload("Flavors").Where("id = ?", id).First(&dish).Error
	if err != nil {
		return nil, err
	}
	return &dish, nil
}

func (d *DishDao) Status(ctx context.Context, id uint64, status int) error {
	err := d.db.WithContext(ctx).Model(&model.Dish{Id: id}).Update("status", status).Error
	return err
}

func (d *DishDao) PageQuery(ctx context.Context, dto *request.DishPageQueryDTO) (*common.PageResult, error) {
	var pageResult common.PageResult
	var dishList []response.DishPageVo
	query := d.db.WithContext(ctx).Model(&model.Dish{})
	if dto.Name != "" {
		query = query.Where("name LIKE", "%"+dto.Name+"%")
	}
	if dto.Status != 0 {
		query = query.Where("status = ?", dto.Status)
	}
	if dto.CategoryId != 0 {
		query = query.Where("category_id = ?", dto.CategoryId)
	}
	if err := query.Count(&pageResult.Total).Error; err != nil {
		return nil, err
	}
	if err := query.Scopes(pageResult.Paginate(&dto.Page, &dto.PageSize)).Select("dish.*,c.name as category_name").Joins("LEFT OUTER JOIN category c ON C.id = dish.category_id").Order("dish_create_time desc").Scan(&dishList).Error; err != nil {
		return nil, err
	}
	pageResult.Records = dishList
	return &pageResult, nil
}

func (d *DishDao) List(ctx context.Context, categoryId uint64) ([]model.Dish, error) {
	var dishList []model.Dish
	err := d.db.WithContext(ctx).Where("category_id = ?", categoryId).Find(&dishList).Error
	return dishList, err
}

func NewDishRepo(db *gorm.DB) repository.DishRepo {
	return &DishDao{db: db}
}
