package model

type Seller struct {
	ID           int64  `json:"id" gorm:"column:id"`
	Name         string `json:"name" gorm:"column:nama"`
	Hp           string `json:"hp" gorm:"column:hp"`
	Address      string `json:"address" gorm:"column:address"`
	Status       int    `json:"status" gorm:"column:status"`
	LastUpdateBy string `json:"last_update_by" gorm:"column:last_update_by"`
	LastUpdate   int64  `json:"last_update" gorm:"column:last_update"`
}

// TableName ...
func (t *Seller) TableName() string {
	return "public.sellers"
}
