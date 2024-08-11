package models

import "gorm.io/gorm"

type Type struct {
	gorm.Model
	Nama string `json:"nama"`
	Foto string `json:"foto"`
}
