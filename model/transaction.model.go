package model

import "com.ddabadi.antarbarang/enumerate"

type Transaksi struct {
	ID                     int64                     `json:"id"`
	TransaksiDate          int64                     `json:"transaksiDate"`
	TransaksiDateStr       string                    `json:"transaksiDateStr"`
	TanggalRequestAntar    int64                     `json:"tanggalRequestAntar"`
	TanggalRequestAntarStr string                    `json:"tanggalRequestAntarStr"`
	JamRequestAntar        string                    `json:"jamRequestAntar"`
	NamaProduct            string                    `json:"namaProduct"`
	CoordinateTujuan       string                    `json:"coordinateTujuan"`
	Keterangan             string                    `json:"keterangan"`
	PhotoAmbil             string                    `json:"photoAmbil"`
	TanggalAmbil           int64                     `json:"tanggalAmbil"`
	TanggalAmbilStr        string                    `json:"jamAmbil"`
	PhotoSampai            string                    `json:"photoSampai"`
	TanggalSampai          int64                     `json:"tanggalSampai"`
	TanggalSampaiStr       string                    `json:"tanggalSampaiStr" `
	IdSeller               int64                     `json:"idSeller"`
	SellerName             string                    `json:"sellerName"`
	IdDriver               int64                     `json:"idDriver"`
	DriverName             string                    `json:"driverName"`
	IdCustomer             int64                     `json:"idCustomer"`
	CustomerName           string                    `json:"customerName"`
	IdAdmin                int64                     `json:"idAdmin"`
	Status                 enumerate.StatusTransaksi `json:"status"`
	LastUpdateBy           string                    `json:"lastUpdateBy"`
	LastUpdate             int64                     `json:"lastUpdate"`
	LastUpdateStr          string                    `json:"lastUpdateStr"`
}

// TableName ...
// func (t *Transaksi) TableName() string {
// 	return "public.transaction_pickup"
// }
