package model

import "com.ddabadi.antarbarang/enumerate"

type Seller struct {
	ID                int64                  `json:"id"`
	Kode              string                 `json:"kode"`
	Password          string                 `json:"password"`
	Nama              string                 `json:"nama"`
	Hp                string                 `json:"hp"`
	Alamat            string                 `json:"alamat"`
	Status            enumerate.StatusRecord `json:"status"`
	LastUpdateBy      string                 `json:"lastUpdateBy"`
	LastUpdate        string                 `json:"lastUpdate"`
	RegionalId        int64                  `json:"regionalId"`
	RegionalGroupName string                 `json:"regionalGroupName"`
	RegionalName      string                 `json:"regionalName"`
}

// TableName ...
func (t *Seller) TableName() string {
	return "public.sellers"
}
