package model

import "com.ddabadi.antarbarang/enumerate"

type Regional struct {
	ID                int64                  `json:"id"`
	Nama              string                 `json:"nama"`
	RegionalGroupId   int64                  `json:"regionalGroupId"`
	RegionalGroupName string                 `json:"regionalGroupName"`
	Status            enumerate.StatusRecord `json:"status"`
	LastUpdateBy      string                 `json:"lastUpdateBy"`
	LastUpdate        string                 `json:"lastUpdate"`
}

// TableName ...
func (t *Regional) TableName() string {
	return "public.regional"
}
