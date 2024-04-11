package model

import (
	"gorm.io/gorm"
	"time"
	"wm-take-out/internal/enum"
)

type Dish struct {
	Id          uint64    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Name        string    `json:"name"`
	CategoryId  uint64    `json:"categoryId"`
	Price       float64   `json:"price"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
	CreateUser  uint64    `json:"createUser"`
	UpdateUser  uint64    `json:"updateUser"`
	// 一对多???什么意思
	Flavors []DishFlavor `json:"flavors"`
}

type DishFlavor struct {
	Id     uint64 `json:"id"`
	DishId uint64 `json:"dish_id"`
	Name   string `json:"name"`
	Value  string `json:"value"`
}

func (d *Dish) BeforeCreate(tx *gorm.DB) error {
	d.CreateTime = time.Now()
	d.UpdateTime = time.Now()
	value := tx.Statement.Context.Value(enum.CurrentId)
	if uid, ok := value.(uint64); ok {
		d.CreateUser = uid
		d.UpdateUser = uid
	}
	return nil
}

func (d *Dish) BeforeUpdate(tx *gorm.DB) error {
	d.UpdateTime = time.Now()
	value := tx.Statement.Context.Value(enum.CurrentId)
	if uid, ok := value.(uint64); ok {
		d.UpdateUser = uid
	}
	return nil
}

func (d *Dish) TableName() string {
	return "dish"
}

func (d *DishFlavor) TableName() string {
	return "dish flavor"
}
