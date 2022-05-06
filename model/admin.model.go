package model

import "com.ddabadi.antarbarang/enumerate"

type Admin struct {
	ID            int64                  `json:"id"`
	Kode          string                 `json:"kode"`
	Nama          string                 `json:"nama"`
	Password      string                 `json:"password"`
	Status        enumerate.StatusRecord `json:"status"`
	LastUpdateBy  string                 `json:"last_update_by"`
	LastUpdate    int64                  `json:"last_update"`
	LastUpdateStr string                 `json:"last_update_str"`
	RoleID        int64                  `json:"roleId"`
}

// TableName ...
func (t *Admin) TableName() string {
	return "public.admins"
}
