package model

import "com.ddabadi.antarbarang/enumerate"

type Customer struct {
	ID            int64                  `json:"id"`
	SellerId      int64                  `json:"sellerId" `
	Nama          string                 `json:"nama"`
	Hp            string                 `json:"hp"`
	Address       string                 `json:"address"`
	Coordinate    string                 `json:"coordinate"`
	Status        enumerate.StatusRecord `json:"status"`
	LastUpdateBy  string                 `json:"lastUpdateBy"`
	LastUpdate    int64                  `json:"lastUpdate"`
	LastUpdateStr string                 `json:"lastUpdateStr"`
}
