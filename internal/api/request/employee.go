package request

type EmployeeLogin struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type EmployeeEditPassWord struct {
	EmpId       uint64 `json:"empId"`
	NewPassword string `json:"newPassword" binding:"required"`
	OldPassword string `json:"oldPassword" binding:"required"`
}

type EmployeeDTO struct {
	Id       uint64 `json:"id"`
	IdNumber string `json:"id_number" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Sex      string `json:"sex" binding:"required"`
	UserName string `json:"username" binding:"required"`
}

type EmployeePageQueryDTO struct {
	Name     string `json:"name"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
}
