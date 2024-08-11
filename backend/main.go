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
	r.POST("/restaurants", controller.CreateRestaurant)
	r.GET("/restaurants", controller.GetRestaurant)
	r.PUT("/restaurants/:id", controller.UpdateRestaurant)
	r.DELETE("/restaurants/:id", controller.DeleteRestaurant)

	//type
	r.POST("/types", controller.CreateType)
	r.GET("/types", controller.GetType)
	r.PUT("/types/:id", controller.UpdateType)
	r.DELETE("/types/:id", controller.DeleteType)
	r.Run(":8080")
}
