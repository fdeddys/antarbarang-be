package model

import (
	"com.ddabadi.antarbarang/enumerate"
)

type Transaksi struct {
	ID                     int64                     `json:"id"`
	TransaksiDate          string                    `json:"transaksiDate"`
	TransaksiDateStr       string                    `json:"transaksiDateStr"`
	TanggalRequestAntar    string                    `json:"tanggalRequestAntar"`
	TanggalRequestAntarStr string                    `json:"tanggalRequestAntarStr"`
	TanggalOnProccess      string                    `json:"tanggalOnProccess"`
	JamRequestAntar        string                    `json:"jamRequestAntar"`
	NamaProduct            string                    `json:"namaProduct"`
	CoordinateTujuan       string                    `json:"coordinateTujuan"`
	Keterangan             string                    `json:"keterangan"`
	PhotoAmbil             string                    `json:"photoAmbil"`
	TanggalAmbil           string                    `json:"tanggalAmbil"`
	TanggalAmbilStr        string                    `json:"tanggalAmbilStr"`
	PhotoSampai            string                    `json:"photoSampai"`
	TanggalSampai          string                    `json:"tanggalSampai"`
	TanggalSampaiStr       string                    `json:"tanggalSampaiStr" `
	IdSeller               int64                     `json:"idSeller"`
	SellerName             string                    `json:"sellerName"`
	SellerAddress          string                    `json:"sellerAddress"`
	SellerHp               string                    `json:"sellerHp"`
	IdDriver               int64                     `json:"idDriver"`
	DriverName             string                    `json:"driverName"`
	IdCustomer             int64                     `json:"idCustomer"`
	CustomerName           string                    `json:"customerName"`
	CustomerAddress        string                    `json:"customerAddress"`
	CustomerHp             string                    `json:"customerHp"`
	IdAdmin                int64                     `json:"idAdmin"`
	Status                 enumerate.StatusTransaksi `json:"status"`
	StatusName             string                    `json:"statusName"`
	LastUpdateBy           string                    `json:"lastUpdateBy"`
	LastUpdate             string                    `json:"lastUpdate"`
	LastUpdateStr          string                    `json:"lastUpdateStr"`
	RegionalSeller         int64                     `json:"regionalSeller"`
	RegionalCustomer       int64                     `json:"regionalCustomer"`
	RegionalGroupSeller    string                    `json:"regionalGroupSeller"`
	RegionalGroupCustomer  string                    `json:"regionalGroupCustomer"`
	CustLng                string                    `json:"custLng"`
	CustLat                string                    `json:"custLat"`
	Biaya                  int64                     `json:"biaya"`
}

// TableName ...
// func (t *Transaksi) TableName() string {
// 	return "public.transaction_pickup"
// }
