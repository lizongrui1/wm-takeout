package service

import (
	"context"
	"log"
	"strconv"
	"wm-take-out/common"
	"wm-take-out/internal/api/request"
	"wm-take-out/internal/enum"
	"wm-take-out/internal/model"
	"wm-take-out/internal/repository"
)

type CategoryService interface {
	UpdateCategory(ctx context.Context, dto request.CategoryDTO) error
	PageQuery(ctx context.Context, dto request.CategoryPageQueryDTO) (*common.PageResult, error)
	SetStatus(ctx context.Context, id uint64, status int) error
	AddCategory(ctx context.Context, dto request.CategoryDTO) error
	DeleteCategory(ctx context.Context, id uint64) error
	List(ctx context.Context) error
}

type CategorySe struct {
	repo repository.CategoryRepo
}

func (cs *CategorySe) UpdateCategory(ctx context.Context, dto request.CategoryDTO) error {
	cate, err := strconv.Atoi(dto.Cate)
	if err != nil {
		log.Printf("无法将 Cate 转换为整数: %v", err)
		return err
	}
	sort, err := strconv.Atoi(dto.Sort)
	if err != nil {
		log.Printf("无法将 Sort 转换为整数: %v", err)
		return err
	}
	err = cs.repo.Update(ctx, model.Category{
		Id:   dto.Id,
		Name: dto.Name,
		Sort: sort,
		Type: cate,
	})
	return err
}

func (cs *CategorySe) PageQuery(ctx context.Context, dto request.CategoryPageQueryDTO) (*common.PageResult, error) {

}

func (cs *CategorySe) SetStatus(ctx context.Context, id uint64, status int) error {
	err := cs.repo.SetStatus(ctx, model.Category{
		Id:     id,
		Status: status,
	})
	return err
}

func (cs *CategorySe) AddCategory(ctx context.Context, dto request.CategoryDTO) error {
	cate, err := strconv.Atoi(dto.Cate)
	if err != nil {
		log.Printf("无法将 Cate 转换为整数: %v", err)
		return err
	}
	sort, err := strconv.Atoi(dto.Sort)
	if err != nil {
		log.Printf("无法将 Sort 转换为整数: %v", err)
		return err
	}
	err = cs.repo.InsertCg(ctx, model.Category{
		Status: enum.ENABLE,
		Name:   dto.Name,
		Sort:   sort,
		Type:   cate,
	})
	return err
}

func (cs *CategorySe) DeleteCategory(ctx context.Context, id uint64) error {
	err := cs.repo.DeleteById(ctx, id)
	return err
}

func (cs *CategorySe) List(ctx context.Context) error {

}
