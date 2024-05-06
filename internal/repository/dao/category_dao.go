package dao

import (
	"context"
	"gorm.io/gorm"
	"wm-take-out/common"
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
	err := c.db.WithContext(ctx).Delete(&model.Category{}, id).Error
	return err
}

func (c *Category) Update(ctx context.Context, category model.Category) error {
	err := c.db.WithContext(ctx).Model(&category).Updates(&category).Error
	return err
}

func (c *Category) PageQuery(ctx context.Context, dto request.CategoryPageQueryDTO) (*common.PageResult, error) {
	var pageResult common.PageResult
	var categoryList []model.Category
	query := c.db.WithContext(ctx).Model(&model.Category{})
	if dto.Name != "" {
		query = query.Where("name like ?", "%"+dto.Name+"%")
	}
	if dto.Cate != 0 {
		query = query.Where("type = ?", dto.Cate)
	}
	if err := query.Count(&pageResult.Total).Error; err != nil {
		return nil, err
	}
	err := query.Scopes(pageResult.Paginate(&dto.Page, &dto.PageSize)).Order("create_time desc").Find(&categoryList).Error
	pageResult.Records = categoryList
	return &pageResult, err
}

func (c *Category) SetStatus(ctx context.Context, category model.Category) error {
	err := c.db.WithContext(ctx).Model(&category).Update("status", category.Status).Error
	return err
}
func (c *Category) List(ctx context.Context, cate int) ([]model.Category, error) {
	var categoryList []model.Category
	err := c.db.WithContext(ctx).Where("type = ?", cate).Order("sort asc").Order("create_time desc").Find(&categoryList).Error
	return categoryList, err
}

func NewCategoryDao(db *gorm.DB) repository.CategoryRepo {
	return &Category{db: db}
}
