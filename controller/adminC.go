package controller

import (
	"net/http"
	"tmi/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//Get
func GetAllAdmin(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var adm []models.Admin

	db.Find(&adm)

	c.JSON(http.StatusOK, gin.H{"data": adm})
}

//Create
func CreateAdmin(c *gin.Context) {
	var input models.Admin

	//JSON Check
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create Admin
	adm := models.Admin{
		Username: input.Username, Password: input.Password, Role: input.Role}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&adm)

	c.JSON(http.StatusOK, gin.H{"data": adm})
}

//GetByID
func GetAdminId(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var adm models.Admin

	if err := db.Where("id = ?", c.Param("id")).First(&adm).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Record Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": adm})
}

//Update
func UpdateAdmin(c *gin.Context) {
	var input models.Admin
	var adm models.Admin
	var updatedInput models.Admin
	db := c.MustGet("db").(*gorm.DB)

	// JSON Check
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", c.Param("id")).First(&adm).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Record Not Found"})
		return
	}
	// Update Data Hasil Input
	updatedInput.Username = input.Username
	updatedInput.Password = input.Password
	updatedInput.Role = input.Role

	db.Model(&adm).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": adm})
}

// Delete
func DeleteAdmin(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var adm models.Admin

	if err := db.Where("id = ?", c.Param("id")).First(&adm).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Record Not Found"})
		return
	}

	db.Delete(&adm)
	c.JSON(http.StatusOK, gin.H{"data": "data was successfully deleted"})
}
