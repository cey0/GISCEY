package models

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	Name      string  `json:"name"`
	TypeID    uint    `json:"type_id"`
	Deskripsi string  `json:"deskripsi"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Type      Type    `gorm:"foreignKey:TypeID"`
}