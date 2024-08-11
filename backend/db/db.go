package db

import (
	"GISCEY/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnDB() {
	db, err := gorm.Open(sqlite.Open("GIS.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database", err)
		return
	}
	db.AutoMigrate(&models.Restaurant{})
	db.AutoMigrate(&models.Type{})

	DB = db

}
