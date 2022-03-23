package model

import "com.ddabadi.antarbarang/enumerate"

type Transaksi struct {
	ID                  int64                     `json:"id" gorm:"column:id"`
	TransaksiDate       int64                     `json:"transaksi_date" gorm:"column:transaksiDate"`
	TanggalRequestAntar int64                     `json:"tanggalRequestAntar" gorm:"column:tanggal_request_antar"`
	JamRequestAntar     string                    `json:"jamRequestAntar" gorm:"column:jam_request_antar"`
	NamaProduct         string                    `json:"namaProduct" gorm:"column:nama_product"`
	CoordinateTujuan    string                    `json:"coordinateTujuan" gorm:"column:coordinate_tujuan"`
	Keterangan          string                    `json:"keterangan" gorm:"column:keterangan"`
	TanggalAmbilStr     string                    `json:"jamAmbil" gorm:"column:jam_ambil"`
	TanggalAmbil        int64                     `json:"tanggalAmbil" gorm:"column:tanggal_ambil"`
	PhotoAmbil          string                    `json:"photoAmbil" gorm:"column:photo_ambil"`
	PhotoSampai         string                    `json:"photoSampai" gorm:"column:photo_sampai"`
	JamSampai           string                    `json:"jamSampai" gorm:"column:jam_sampai"`
	TanggalSampai       int8                      `json:"tanggalSampai" gorm:"column:tanggal_sampai"`
	IdSeller            int64                     `json:"idSeller" gorm:"column:id_seller"`
	IdDriver            int64                     `json:"idDriver" gorm:"column:id_driver"`
	IdCustomer          int64                     `json:"idCustomer" gorm:"column:id_customer"`
	IdAdmin             int64                     `json:"idAdmin" gorm:"column:id_admin"`
	Status              enumerate.StatusTransaksi `json:"status" gorm:"column:status"`
	LastUpdateBy        string                    `json:"last_update_by" gorm:"column:last_update_by"`
	LastUpdate          int64                     `json:"last_update" gorm:"column:last_update"`
}

// TableName ...
func (t *Transaksi) TableName() string {
	return "public.transaction_pickup"
}
