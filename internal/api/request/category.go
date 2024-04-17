package request

type CategoryDTO struct {
	Id   uint64 `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Name string `json:"name"`
	Sort string `json:"sort"`
	Cate string `json:"type"`
}

type CategoryPageQueryDTO struct {
	Name     string `json:"name"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	Cate     int    `json:"type"`
}
