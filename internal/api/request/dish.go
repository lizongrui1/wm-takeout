package request

import "wm-take-out/internal/model"

type DishDTO struct {
	Id          uint64             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Image       string             `json:"image"`
	Price       uint64             `json:"price"`
	Status      int                `json:"status"`
	CategoryId  uint64             `json:"categoryId"`
	Flavors     []model.DishFlavor `json:"flavors"`
}

type DishPageQueryDto struct {
	Name     string `json:"name"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
}
