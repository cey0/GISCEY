package db

import (
	"GISCEY/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnDB() {
	db, err := gorm.Open(mysql.Open("root:020906@tcp(127.0.0.1:3306)/GIS"), &gorm.Config{}) // Ganti sqlite dengan mariadb
	if err != nil {
		log.Fatal("failed to connect to database", err)
		return
	}
	db.AutoMigrate(&models.Restaurant{})
	db.AutoMigrate(&models.Type{})
	db.AutoMigrate(&models.User{})

	DB = db

}
