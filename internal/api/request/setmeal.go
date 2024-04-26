package request

import "wm-take-out/internal/model"

type SetMealDTO struct {
	Id           uint64              `json:"id"`            // 主键id
	CategoryId   uint64              `json:"categoryId"`    // 分类id
	Name         string              `json:"name"`          // 套餐名称
	Price        string              `json:"price"`         // 套餐单价
	Status       int                 `json:"status"`        // 套餐状态
	Description  string              `json:"description"`   // 套餐描述
	Image        string              `json:"image"`         // 套餐图片
	SetMealDishs []model.SetMealDish `json:"setmealDishes"` // 套餐菜品关系，一个套餐可以包含多个菜品，且一个菜品可以属于多个套餐，所以它们之间形成了多对多的关系。在数据库中直接存储多对多关系是不可行的，因此需要一个额外的表来表示这种关系。
}

type SetMealPageQueryDTO struct {
	Page       int    `json:"page"`       // 分页查询的页数
	PageSize   int    `json:"pageSize"`   // 分页查询的页容量
	Name       string `json:"name"`       // 分页查询的name
	CategoryId uint64 `json:"categoryId"` // 分类ID
	Status     int    `json:"status"`     // 套餐起售状态
}
