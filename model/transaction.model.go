package model

import "com.ddabadi.antarbarang/enumerate"

type Transaksi struct {
	ID                  int64                     `json:"id"`
	TransaksiDate       int64                     `json:"transaksiDate"`
	TanggalRequestAntar int64                     `json:"tanggalRequestAntar"`
	JamRequestAntar     string                    `json:"jamRequestAntar"`
	NamaProduct         string                    `json:"namaProduct"`
	CoordinateTujuan    string                    `json:"coordinateTujuan"`
	Keterangan          string                    `json:"keterangan"`
	PhotoAmbil          string                    `json:"photoAmbil"`
	TanggalAmbil        int64                     `json:"tanggalAmbil"`
	TanggalAmbilStr     string                    `json:"jamAmbil"`
	PhotoSampai         string                    `json:"photoSampai"`
	TanggalSampai       int64                     `json:"tanggalSampai"`
	TanggalSampaiStr    string                    `json:"tanggalSampaiStr" `
	IdSeller            int64                     `json:"idSeller"`
	IdDriver            int64                     `json:"idDriver"`
	IdCustomer          int64                     `json:"idCustomer"`
	IdAdmin             int64                     `json:"idAdmin"`
	Status              enumerate.StatusTransaksi `json:"status"`
	LastUpdateBy        string                    `json:"last_update_by"`
	LastUpdate          int64                     `json:"last_update"`
}

// TableName ...
func (t *Transaksi) TableName() string {
	return "public.transaction_pickup"
}
