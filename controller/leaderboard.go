package controller

import (
	"net/http"
	"tmi/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//Get
func GetAllLeaderboard(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var ttgn []models.Leaderboard

	db.Find(&ttgn)

	c.JSON(http.StatusOK, gin.H{"data": ttgn})
}

//Create
func CreateLeaderboard(c *gin.Context) {
	var input models.Leaderboard

	//JSON Check
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create Leaderboard
	ttgn := models.Leaderboard{
		JudulLeaderboard: input.JudulLeaderboard, Tema: input.Tema, Latar: input.Latar, UrlVideoLeaderboard: input.UrlVideoLeaderboard, TaskLeaderboard: input.TaskLeaderboard, UjianLeaderboard: input.UjianLeaderboard, SkorLeaderboard: input.SkorLeaderboard, IDAdmin: input.IDAdmin}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&ttgn)

	c.JSON(http.StatusOK, gin.H{"data": ttgn})
}

//GetByID
func GetLeaderboardId(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var ttgn models.Leaderboard

	if err := db.Where("id = ?", c.Param("id")).First(&ttgn).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Record Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ttgn})
}

//Update
func UpdateLeaderboard(c *gin.Context) {
	var input models.Leaderboard
	var ttgn models.Leaderboard
	var updatedInput models.Leaderboard
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
	updatedInput.JudulLeaderboard = input.JudulLeaderboard
	updatedInput.Tema = input.Tema
	updatedInput.Latar = input.Latar
	updatedInput.UrlVideoLeaderboard = input.UrlVideoLeaderboard
	updatedInput.TaskLeaderboard = input.TaskLeaderboard
	updatedInput.UjianLeaderboard = input.UjianLeaderboard
	updatedInput.SkorLeaderboard = input.SkorLeaderboard
	updatedInput.IDAdmin = input.IDAdmin

	db.Model(&ttgn).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": ttgn})
}

// Delete
func DeleteLeaderboard(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var ttgn models.Leaderboard

	if err := db.Where("id = ?", c.Param("id")).First(&ttgn).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Record Not Found"})
		return
	}

	db.Delete(&ttgn)
	c.JSON(http.StatusOK, gin.H{"data": "data was successfully deleted"})
}
