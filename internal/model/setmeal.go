package model

import (
	"gorm.io/gorm"
	"time"
	"wm-take-out/internal/enum"
)

type SetMeal struct {
	Id          uint64    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	CategoryId  uint64    `json:"category_id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Status      int       `json:"status"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
	CreateUser  uint64    `json:"create_user"`
	UpdateUser  uint64    `json:"update_user"`
}

func (s *SetMeal) BeforeCreate(tx *gorm.DB) error {
	s.CreateTime = time.Now()
	s.UpdateTime = time.Now()
	value := tx.Statement.Context.Value(enum.CurrentId)
	if uid, ok := value.(uint64); ok {
		s.CreateUser = uid
		s.UpdateUser = uid
	}
	return nil
}

func (s *SetMeal) BeforeUpdate(tx *gorm.DB) error {
	s.UpdateTime = time.Now()
	value := tx.Statement.Context.Value(enum.CurrentId)
	if uid, ok := value.(uint64); ok {
		s.UpdateUser = uid
	}
	return nil
}

func (s *SetMeal) TableName() string {
	return "setmeal"
}
