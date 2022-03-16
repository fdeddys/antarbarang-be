package model

type Urut struct {
	ID         int64  `json:"id" gorm:"column:id"`
	Prefix     string `json:"prefix" gorm:"column:prefix"`
	Keterangan string `json:"keterangan" gorm:"column:keterangan"`
	NoTerakhir int64  `json:"no_terakhir" gorm:"column:no_terakhir"`
}

// TableName ...
func (t *Urut) TableName() string {
	return "public.uruts"
}
