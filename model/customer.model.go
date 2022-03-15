package model

import "com.ddabadi.antarbarang/enumerate"

type Customer struct {
	ID           int64                  `json:"id" gorm:"column:id"`
	Nama         string                 `json:"nama" gorm:"column:name"`
	Hp           string                 `json:"hp" gorm:"column:hp"`
	Address      string                 `json:"address" gorm:"column:address"`
	Coordinate   string                 `json:"coordinate" gorm:"column:coordinate"`
	Status       enumerate.StatusRecord `json:"status" gorm:"column:status"`
	LastUpdateBy string                 `json:"last_update_by" gorm:"column:last_update_by"`
	LastUpdate   int64                  `json:"last_update" gorm:"column:last_update"`
}

// TableName ...
func (t *Customer) TableName() string {
	return "public.customers"
}
