package model

import (
	"gorm.io/gorm"
	"time"
)

type employee struct {
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

func (e employee) BeforeCreate(tx *gorm.DB) error {
	
}
