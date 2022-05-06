package model

import "time"

// RoleMenu ...
type RoleMenu struct {
	RoleID       int       `json:"roleId"`
	MenuID       int       `json:"menuId"`
	Status       int       `json:"status"`
	LastUpdateBy string    `json:"lastUpdateBy"`
	LastUpdate   time.Time `json:"lastUpdate"`
}
