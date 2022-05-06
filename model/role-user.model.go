package model

import "time"

// Role ...
type RoleUser struct {
	RoleID       int64     `json:"roleId"`
	UserID       int64     `json:"userId"`
	Status       int       `json:"status"`
	LastUpdateBy string    `json:"lastUpdateBy"`
	LastUpdate   time.Time `json:"lastUpdate"`
}
