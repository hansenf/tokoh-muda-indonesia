package controller

import (
	"net/http"
	"tmi/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//Get
func GetAllTantangan(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var ttgn []models.Tantangan

	db.Find(&ttgn)

	c.JSON(http.StatusOK, gin.H{"data": ttgn})
}

//Create
func CreateTantangan(c *gin.Context) {
	var input models.Tantangan

	//JSON Check
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create Tantangan
	ttgn := models.Tantangan{
		JudulTantangan: input.JudulTantangan, Tema: input.Tema, Latar: input.Latar, UrlVideoTantangan: input.UrlVideoTantangan, TaskTantangan: input.TaskTantangan, UjianTantangan: input.UjianTantangan, SkorTantangan: input.SkorTantangan, IDAdmin: input.IDAdmin}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&ttgn)

	c.JSON(http.StatusOK, gin.H{"data": ttgn})
}

//GetByID
func GetTantanganId(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var ttgn models.Tantangan

	if err := db.Where("id = ?", c.Param("id")).First(&ttgn).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Record Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ttgn})
}

//Update
func UpdateTantangan(c *gin.Context) {
	var input models.Tantangan
	var ttgn models.Tantangan
	var updatedInput models.Tantangan
	db := c.MustGet("db").(*gorm.DB)

	// JSON Check
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", c.Param("id")).First(&ttgn).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Record Not Found"})
		return
	}
	// Update Data Hasil Input
	updatedInput.JudulTantangan = input.JudulTantangan
	updatedInput.Tema = input.Tema
	updatedInput.Latar = input.Latar
	updatedInput.UrlVideoTantangan = input.UrlVideoTantangan
	updatedInput.TaskTantangan = input.TaskTantangan
	updatedInput.UjianTantangan = input.UjianTantangan
	updatedInput.SkorTantangan = input.SkorTantangan
	updatedInput.IDAdmin = input.IDAdmin

	db.Model(&ttgn).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": ttgn})
}

// Delete
func DeleteTantangan(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var ttgn models.Tantangan

	if err := db.Where("id = ?", c.Param("id")).First(&ttgn).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Record Not Found"})
		return
	}

	db.Delete(&ttgn)
	c.JSON(http.StatusOK, gin.H{"data": "data was successfully deleted"})
}
