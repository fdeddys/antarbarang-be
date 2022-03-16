package model

import "com.ddabadi.antarbarang/enumerate"

type Seller struct {
	ID           int64                  `json:"id" gorm:"column:id"`
	Kode         string                 `json:"kode" gorm:"column:kode"`
	Password     string                 `json:"password" gorm:"column:password"`
	Nama         string                 `json:"nama" gorm:"column:nama"`
	Hp           string                 `json:"hp" gorm:"column:hp"`
	Alamat       string                 `json:"alamat" gorm:"column:alamat"`
	Status       enumerate.StatusRecord `json:"status" gorm:"column:status"`
	LastUpdateBy string                 `json:"last_update_by" gorm:"column:last_update_by"`
	LastUpdate   int64                  `json:"last_update" gorm:"column:last_update"`
}

// TableName ...
func (t *Seller) TableName() string {
	return "public.sellers"
}
