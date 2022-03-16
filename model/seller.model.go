package model

import "com.ddabadi.antarbarang/enumerate"

type Seller struct {
	ID            int64                  `json:"id"`
	Kode          string                 `json:"kode"`
	Password      string                 `json:"password"`
	Nama          string                 `json:"nama"`
	Hp            string                 `json:"hp"`
	Alamat        string                 `json:"alamat"`
	Status        enumerate.StatusRecord `json:"status"`
	LastUpdateBy  string                 `json:"last_update_by"`
	LastUpdate    int64                  `json:"last_update"`
	LastUpdateStr string                 `json:"last_update_str"`
}

// TableName ...
func (t *Seller) TableName() string {
	return "public.sellers"
}
