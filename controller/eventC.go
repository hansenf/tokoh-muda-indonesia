package controller

import (
	"net/http"
	"tmi/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//Get
func GetAllEvent(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var evnt []models.Event

	db.Find(&evnt)

	c.JSON(http.StatusOK, gin.H{"data": evnt})
}

//Create
func CreateEvent(c *gin.Context) {
	var input models.Event

	//JSON Check
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create Event
	evnt := models.Event{
		JudulEvent: input.JudulEvent, DeskripsiEvent: input.DeskripsiEvent, KriteriaEvent: input.KriteriaEvent, TanggalEvent: input.TanggalEvent, IDAdmin: input.IDAdmin}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&evnt)

	c.JSON(http.StatusOK, gin.H{"data": evnt})
}

//GetByID
func GetEventId(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var evnt models.Event

	if err := db.Where("id = ?", c.Param("id")).First(&evnt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Record Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": evnt})
}

//Update
func UpdateEvent(c *gin.Context) {
	var input models.Event
	var evnt models.Event
	var updatedInput models.Event
	db := c.MustGet("db").(*gorm.DB)

	// JSON Check
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", c.Param("id")).First(&evnt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Record Not Found"})
		return
	}
	// Update Data Hasil Input
	updatedInput.JudulEvent = input.JudulEvent
	updatedInput.DeskripsiEvent = input.DeskripsiEvent
	updatedInput.KriteriaEvent = input.KriteriaEvent
	updatedInput.TanggalEvent = input.TanggalEvent
	updatedInput.IDAdmin = input.IDAdmin

	db.Model(&evnt).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": evnt})
}

// Delete
func DeleteEvent(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var evnt models.Event

	if err := db.Where("id = ?", c.Param("id")).First(&evnt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Record Not Found"})
		return
	}

	db.Delete(&evnt)
	c.JSON(http.StatusOK, gin.H{"data": "data was successfully deleted"})
}
