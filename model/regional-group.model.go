package model

import (
	"com.ddabadi.antarbarang/enumerate"
)

type RegionalGroup struct {
	ID           int64                  `json:"id"`
	Nama         string                 `json:"nama"`
	Status       enumerate.StatusRecord `json:"status"`
	LastUpdateBy string                 `json:"lastUpdateBy"`
	LastUpdate   string                 `json:"lastUpdate"`
}

// TableName ...
func (t *RegionalGroup) TableName() string {
	return "public.regional_group"
}
