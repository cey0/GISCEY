package main

import (
	"GISCEY/controller"
	"GISCEY/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db.ConnDB()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:5173"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}

	r.Use(cors.New(corsConfig))

	r.Static("/fotoTypes", "../fotoTypes")

	//restaurant
	r.POST("/api/restaurants", controller.CreateRestaurant)
	r.GET("/api/restaurants", controller.GetRestaurant)
	r.PUT("/api/restaurants/:id", controller.UpdateRestaurant)
	r.DELETE("/api/restaurants/:id", controller.DeleteRestaurant)

	//type
	r.POST("/api/types", controller.CreateType)
	r.GET("/api/types", controller.GetType)
	r.PUT("/api/types/:id", controller.UpdateType)
	r.DELETE("/api/types/:id", controller.DeleteType)
	r.Run(":8080")
}
