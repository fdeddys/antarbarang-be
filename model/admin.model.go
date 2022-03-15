package model

type Admin struct {
	ID           int64  `json:"id" gorm:"column:id"`
	Nama         string `json:"nama" gorm:"column:nama"`
	Status       int    `json:"status" gorm:"column:status"`
	LastUpdateBy string `json:"last_update_by" gorm:"column:last_update_by"`
	LastUpdate   int64  `json:"last_update" gorm:"column:last_update"`
}

// TableName ...
func (t *Admin) TableName() string {
	return "public.admins"
}
