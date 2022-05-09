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
	LastUpdateBy  string                 `json:"lastUpdateBy"`
	LastUpdate    int64                  `json:"lastUpdate"`
	LastUpdateStr string                 `json:"lastUpdateStr"`
}

// // TableName ...
// func (t *Driver) TableName() string {
// 	return "public.drivers"
// }
