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
	var ldbrd []models.Leaderboard

	db.Find(&ldbrd)

	c.JSON(http.StatusOK, gin.H{"data": ldbrd})
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
	ldbrd := models.Leaderboard{
		Ranking: input.Ranking, IDMahasiswa: input.IDMahasiswa, IDTantangan: input.IDTantangan, IDSilabus: input.IDSilabus}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&ldbrd)

	c.JSON(http.StatusOK, gin.H{"data": ldbrd})
}

//GetByID
func GetLeaderboardId(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var ldbrd models.Leaderboard

	if err := db.Where("id = ?", c.Param("id")).First(&ldbrd).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Record Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ldbrd})
}

//Update
func UpdateLeaderboard(c *gin.Context) {
	var input models.Leaderboard
	var ldbrd models.Leaderboard
	var updatedInput models.Leaderboard
	db := c.MustGet("db").(*gorm.DB)

	// JSON Check
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", c.Param("id")).First(&ldbrd).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Record Not Found"})
		return
	}
	// Update Data Hasil Input
	updatedInput.Ranking = input.Ranking
	updatedInput.IDMahasiswa = input.IDMahasiswa
	updatedInput.IDTantangan = input.IDTantangan
	updatedInput.IDSilabus = input.IDSilabus

	db.Model(&ldbrd).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": ldbrd})
}

// Delete
func DeleteLeaderboard(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var ldbrd models.Leaderboard

	if err := db.Where("id = ?", c.Param("id")).First(&ldbrd).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Record Not Found"})
		return
	}

	db.Delete(&ldbrd)
	c.JSON(http.StatusOK, gin.H{"data": "data was successfully deleted"})
}
