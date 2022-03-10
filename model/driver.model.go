package model

type Driver struct {
	ID           int64  `json:"id" gorm:"column:id"`
	Picture      string `json:"picture" gorm:"column:pict"`
	Address      string `json:"address" gorm:"column:address"`
	Hp           string `json:"hp" gorm:"column:hp"`
	Name         string `json:"name" gorm:"column:name"`
	Status       int    `json:"status" gorm:"column:status"`
	LastUpdateBy string `json:"last_update_by" gorm:"column:last_update_by"`
	LastUpdate   string `json:"last_update" gorm:"column:last_update"`
	Code         string `json:"code" gorm:"column:code"`
}

// TableName ...
func (t *Driver) TableName() string {
	return "public.driver"
}
