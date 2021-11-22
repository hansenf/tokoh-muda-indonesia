package controller

import (
	"net/http"
	"tmi/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//Get
func GetAllSilabus(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var slb []models.Silabus

	db.Find(&slb)

	c.JSON(http.StatusOK, gin.H{"data": slb})
}

//Create
func CreateSilabus(c *gin.Context) {
	var input models.Silabus

	//JSON Check
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create Silabus
	slb := models.Silabus{
		JudulSilabus: input.JudulSilabus, Definisi: input.Definisi, FungsiSilabus: input.FungsiSilabus, Deskripsi: input.Deskripsi, UrlVideoSilabus: input.UrlVideoSilabus, TaskSilabus: input.TaskSilabus, UjianSilabus: input.UjianSilabus, SkorSilabus: input.SkorSilabus, IDAdmin: input.IDAdmin}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&slb)

	c.JSON(http.StatusOK, gin.H{"data": slb})
}

//GetByID
func GetSilabusId(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var slb models.Silabus

	if err := db.Where("id = ?", c.Param("id")).First(&slb).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Record Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": slb})
}

//Update
func UpdateSilabus(c *gin.Context) {
	var input models.Silabus
	var slb models.Silabus
	var updatedInput models.Silabus
	db := c.MustGet("db").(*gorm.DB)

	// JSON Check
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", c.Param("id")).First(&slb).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Record Not Found"})
		return
	}
	// Update Data Hasil Input
	updatedInput.JudulSilabus = input.JudulSilabus
	updatedInput.Definisi = input.Definisi
	updatedInput.FungsiSilabus = input.FungsiSilabus
	updatedInput.Deskripsi = input.Deskripsi
	updatedInput.UrlVideoSilabus = input.UrlVideoSilabus
	updatedInput.TaskSilabus = input.TaskSilabus
	updatedInput.UjianSilabus = input.UjianSilabus
	updatedInput.SkorSilabus = input.SkorSilabus
	updatedInput.IDAdmin = input.IDAdmin

	db.Model(&slb).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": slb})
}

// Delete
func DeleteSilabus(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var slb models.Silabus

	if err := db.Where("id = ?", c.Param("id")).First(&slb).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Record Not Found"})
		return
	}

	db.Delete(&slb)
	c.JSON(http.StatusOK, gin.H{"data": "data was successfully deleted"})
}
