package controller

import (
	"GISCEY/db"
	"GISCEY/models"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateType(c *gin.Context) {
	nama := c.PostForm("nama")

	file, err := c.FormFile("foto")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed: " + err.Error()})
		return
	}

	filename := strconv.FormatInt(time.Now().Unix(), 10) + "_" + file.Filename

	filePath := filepath.Join("../fotoTypes", filename)

	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create directory: " + err.Error()})
			return
		}
	}

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file: " + err.Error()})
		return
	}

	res := models.Type{
		Nama: nama,
		Foto: filename,
	}

	if err := db.DB.Create(&res).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": res})
}

func GetType(c *gin.Context) {
	var res []models.Type
	db.DB.Find(&res)
	c.JSON(http.StatusOK, gin.H{"success": res})
}

func UpdateType(c *gin.Context) {
	idStr := c.Param("id")         // Get ID as string
	id, err := strconv.Atoi(idStr) // Convert to integer
	if err != nil {                // Add error checking
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var Type models.Type
	if err := db.DB.First(&Type, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Type not found"})
		return
	}

	// Parse form data
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Form parsing error: " + err.Error()})
		return
	}

	nama := c.PostForm("nama")
	Type.Nama = nama

	file, err := c.FormFile("foto")
	if err == nil {
		filename := strconv.FormatInt(time.Now().Unix(), 10) + "_" + file.Filename
		filePath := filepath.Join("../fotoTypes", filename)

		dir := filepath.Dir(filePath)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create directory: " + err.Error()})
				return
			}
		}

		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file: " + err.Error()})
			return
		}

		Type.Foto = filename
	}

	if err := db.DB.Save(&Type).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, Type)
}

func DeleteType(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	var Type models.Type
	if err := db.DB.First(&Type, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Type not found"})
		return
	}

	if err := db.DB.Delete(&Type).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete type: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Type deleted"})
}
