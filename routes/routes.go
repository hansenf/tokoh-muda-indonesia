package routes

import (
	"tmi/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// middleware
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"admin":  "password",
		"editor": "rahasia",
	}))

	r.GET("/api/mahasiswa", controller.GetAllMahasiswa)
	authorized.POST("/api/mahasiswa", controller.CreateMahasiswa)
	r.GET("/api/mahasiswa/:id", controller.GetMahasiswaId)
	authorized.PUT("/api/mahasiswa/:id", controller.UpdateMahasiswa)
	authorized.DELETE("/api/mahasiswa/:id", controller.DeleteMahasiswa)

	return r
}
