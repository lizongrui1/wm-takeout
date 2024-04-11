package model

import (
	"gorm.io/gorm"
	"time"
	"wm-take-out/internal/enum"
)

type Category struct {
	Id         uint64    `json:"id"`
	Type       int       `json:"type"`
	Name       string    `json:"name"`
	Sort       int       `json:"sort"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	CreateUser uint64    `json:"createUser"`
	UpdateUser uint64    `json:"updateUser"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) error {
	c.CreateTime = time.Now()
	c.UpdateTime = time.Now()
	value := tx.Statement.Context.Value(enum.CurrentId)
	if uid, ok := value.(uint64); ok {
		c.CreateUser = uid
		c.UpdateUser = uid
	}
	return nil
}

func (c *Category) BeforeUpdate(tx *gorm.DB) error {
	c.UpdateTime = time.Now()
	value := tx.Statement.Context.Value(enum.CurrentId)
	if uid, ok := value.(uint64); ok {
		c.UpdateUser = uid
	}
	return nil
}
