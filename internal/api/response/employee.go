package response

import "time"

type EmployeeLogin struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Token    string `json:"token"`
	UserName string `json:"username"`
}

type EmployeeQueryById struct {
	Id         uint64    `json:"id"`
	IdNumber   uint64    `json:"idNumber"`
	Name       string    `json:"name"`
	UserName   string    `json:"username"`
	Password   string    `json:"password"`
	Phone      string    `json:"phone"`
	Sex        string    `json:"sex"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	CreateUser uint64    `json:"createUser"`
	UpdateUser uint64    `json:"updateUser"`
}
