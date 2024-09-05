package controller

import (
	"GISCEY/controller/request"
	"GISCEY/db"
	"GISCEY/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(c *gin.Context) {
	var req request.CreateRestaurantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := models.Restaurant{
		Name:      req.Name,
		TypeID:    req.TypeID,
		Deskripsi: req.Deskripsi,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}
	db.DB.Create(&res)
	c.JSON(http.StatusCreated, gin.H{"success": res})
}

func GetRestaurant(c *gin.Context) {
	var res []models.Restaurant

	db.DB.Preload("Type").Find(&res)
	c.JSON(http.StatusOK, gin.H{"success": res})
}

func UpdateRestaurant(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var restaurant models.Restaurant
	if err := db.DB.First(&restaurant, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
		return
	}

	var input request.UpdateRestaurantRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&restaurant).Updates(input)
	c.JSON(http.StatusOK, restaurant)
}

func DeleteRestaurant(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var restaurant models.Restaurant
	if err := db.DB.First(&restaurant, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
		return
	}

	db.DB.Delete(&restaurant)
	c.JSON(http.StatusOK, gin.H{"success": "Restaurant deleted"})
}
