package model

import "com.ddabadi.antarbarang/enumerate"

type Driver struct {
	ID            int64                  `json:"id"`
	Kode          string                 `json:"kode"`
	Photo         string                 `json:"photo"`
	Alamat        string                 `json:"alamat"`
	Password      string                 `json:"password"`
	Hp            string                 `json:"hp"`
	Nama          string                 `json:"nama"`
	Status        enumerate.StatusRecord `json:"status"`
	LastUpdateBy  string                 `json:"last_update_by"`
	LastUpdate    int64                  `json:"last_update"`
	LastUpdateStr string                 `json:"last_update_str"`
}

// // TableName ...
// func (t *Driver) TableName() string {
// 	return "public.drivers"
// }
