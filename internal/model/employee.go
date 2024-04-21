package model

import (
	"gorm.io/gorm"
	"time"
	"wm-take-out/internal/enum"
)

type Employee struct {
	Id         uint64    `json:"id"`
	UserName   string    `json:"user_name"`
	Name       string    `json:"name"`
	Password   string    `json:"password"`
	Phone      string    `json:"phone"`
	Sex        string    `json:"sex"`
	IdNumber   string    `json:"idNumber"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	CreateUser uint64    `json:"createUser"`
	UpdateUser uint64    `json:"updateUser"`
}

func (e *Employee) BeforeCreate(tx *gorm.DB) error {
	e.CreateTime = time.Now()
	e.UpdateTime = time.Now()
	value := tx.Statement.Context.Value(enum.CurrentId)
	if uid, ok := value.(uint64); ok {
		e.CreateUser = uid
		e.UpdateUser = uid
	}
	return nil
}

func (e *Employee) BeforeUpdate(tx *gorm.DB) error {
	e.UpdateTime = time.Now()
	value := tx.Statement.Context.Value(enum.CurrentId)
	if uid, ok := value.(uint64); ok {
		e.UpdateUser = uid
	}
	return nil
}

func (e *Employee) AfterFind(tx *gorm.DB) error {
	e.CreateTime.Format(time.DateOnly)
	e.CreateTime.Format(time.DateTime)
	return nil
}

func (e *Employee) TableName() string {
	return "employee"
}
