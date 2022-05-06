package model

import "time"

type Parameter struct {
	ID           int64     `json:"id" gorm:"column:id"`
	Nama         string    `json:"name" gorm:"column:nama"`
	Value        string    `json:"value" gorm:"column:value"`
	IsViewable   int8      `json:"IsViewable" gorm:"column:isviewable"`
	LastUpdateBy string    `json:"last_update_by" gorm:"column:last_update_by"`
	LastUpdate   time.Time `json:"last_update" gorm:"column:last_update"`
}
