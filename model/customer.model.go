package model

import "com.ddabadi.antarbarang/enumerate"

type Customer struct {
	ID                int64                  `json:"id"`
	SellerId          int64                  `json:"sellerId" `
	SellerName        string                 `json:"sellerName" `
	Nama              string                 `json:"nama"`
	Hp                string                 `json:"hp"`
	Alamat            string                 `json:"alamat"`
	Coordinate        string                 `json:"coordinate"`
	Status            enumerate.StatusRecord `json:"status"`
	LastUpdateBy      string                 `json:"lastUpdateBy"`
	LastUpdate        string                 `json:"lastUpdate"`
	RegionalId        int64                  `json:"regionalId"`
	RegionalGroupName string                 `json:"regionalGroupName"`
	RegionalName      string                 `json:"regionalName"`
	Lng               string                 `json:"lng"`
	Lat               string                 `json:"lat"`
}
