package model

type SetMealDish struct {
	Id        uint64  `json:"id"`        // 中间表id
	SetmealId uint64  `json:"setmealId"` // 套餐id 存储与之关联的套餐的 ID。这是一个外键，指向套餐表的主键。
	DishId    uint64  `json:"dishId"`    // 菜品id 存储与之关联的菜品的 ID。这是一个外键，指向菜品表的主键。
	Name      string  `json:"name"`      // 菜品名称冗余字段
	Price     float64 `json:"price"`     // 菜品单价冗余字段
	Copies    int     `json:"copies"`    // 菜品数量冗余字段
}

func (e *SetMealDish) TableName() string {
	return "setmeal_dish"
}
