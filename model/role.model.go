package model

import "time"

// Role ...
type Role struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	LastUpdateBy string    `json:"lastUpdateBy"`
	LastUpdate   time.Time `json:"lastUpdate"`
}
