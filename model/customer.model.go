package model

import "com.ddabadi.antarbarang/enumerate"

type Customer struct {
	ID            int64                  `json:"id"`
	SellerId      int64                  `json:"seller_id" `
	Nama          string                 `json:"nama"`
	Hp            string                 `json:"hp"`
	Alamat        string                 `json:"alamat"`
	Coordinate    string                 `json:"coordinate"`
	Status        enumerate.StatusRecord `json:"status"`
	LastUpdateBy  string                 `json:"last_update_by"`
	LastUpdate    int64                  `json:"last_update"`
	LastUpdateStr string                 `json:"last_update_str"`
}
