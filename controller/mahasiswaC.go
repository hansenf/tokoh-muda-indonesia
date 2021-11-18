package controller

import (
	"net/http"
	"tmi/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//Get
func GetAllMahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var mhs []models.Mahasiswa

	db.Find(&mhs)

	c.JSON(http.StatusOK, gin.H{"data": mhs})
}

//Create
func CreateMahasiswa(c *gin.Context) {
	var input models.Mahasiswa

	//JSON Check
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create mahasiswa
	mhs := models.Mahasiswa{
		Username: input.Username, Password: input.Password, Nomorhp: input.Nomorhp,
		Email: input.Email, UrlFoto: input.UrlFoto, NamaLengkap: input.NamaLengkap,
		TanggalLahir: input.TanggalLahir, JenisKelamin: input.JenisKelamin, AsalKampus: input.AsalKampus,
		Nim: input.Nim, Jurusan: input.Jurusan, TahunMasuk: input.TahunMasuk, KotaKabupaten: input.KotaKabupaten,
		IDTantangan: input.IDTantangan, IDSilabus: input.IDSilabus, IDEvent: input.IDEvent}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&mhs)

	c.JSON(http.StatusOK, gin.H{"data": mhs})
}

//GetByID
func GetMahasiswaId(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var mhs models.Mahasiswa

	if err := db.Where("id = ?", c.Param("id")).First(&mhs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Record Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": mhs})
}

//Update
func UpdateMahasiswa(c *gin.Context) {
	var input models.Mahasiswa
	var mhs models.Mahasiswa
	var updatedInput models.Mahasiswa
	db := c.MustGet("db").(*gorm.DB)

	// JSON Check
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", c.Param("id")).First(&mhs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Record Not Found"})
		return
	}
	// Update Data Hasil Input
	updatedInput.Username = input.Username
	updatedInput.Password = input.Password
	updatedInput.Nomorhp = input.Nomorhp
	updatedInput.Email = input.Email
	updatedInput.UrlFoto = input.UrlFoto
	updatedInput.NamaLengkap = input.NamaLengkap
	updatedInput.TanggalLahir = input.TanggalLahir
	updatedInput.JenisKelamin = input.JenisKelamin
	updatedInput.Nim = input.Nim
	updatedInput.Jurusan = input.Jurusan
	updatedInput.TahunMasuk = input.TahunMasuk
	updatedInput.KotaKabupaten = input.KotaKabupaten
	updatedInput.IDTantangan = input.IDTantangan
	updatedInput.IDSilabus = input.IDSilabus
	updatedInput.IDEvent = input.IDEvent

	db.Model(&mhs).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": mhs})
}

// Delete
func DeleteMahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var mhs models.Mahasiswa

	if err := db.Where("id = ?", c.Param("id")).First(&mhs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Record Not Found"})
		return
	}

	db.Delete(&mhs)
	c.JSON(http.StatusOK, gin.H{"data": "data was successfully deleted"})
}
