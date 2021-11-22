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

	// route mahasiswa
	r.GET("/api/mahasiswa", controller.GetAllMahasiswa)
	authorized.POST("/api/mahasiswa", controller.CreateMahasiswa)
	r.GET("/api/mahasiswa/:id", controller.GetMahasiswaId)
	authorized.PUT("/api/mahasiswa/:id", controller.UpdateMahasiswa)
	authorized.DELETE("/api/mahasiswa/:id", controller.DeleteMahasiswa)

	//route admin
	r.GET("/api/admin", controller.GetAllAdmin)
	authorized.POST("/api/admin", controller.CreateAdmin)
	r.GET("/api/admin/:id", controller.GetAdminId)
	authorized.PUT("/api/admin/:id", controller.UpdateAdmin)
	authorized.DELETE("/api/admin/:id", controller.DeleteAdmin)

	//route tantangan
	r.GET("/api/tantangan", controller.GetAllTantangan)
	authorized.POST("/api/tantangan", controller.CreateTantangan)
	r.GET("/api/tantangan/:id", controller.GetTantanganId)
	authorized.PUT("/api/tantangan/:id", controller.UpdateTantangan)
	authorized.DELETE("/api/tantangan/:id", controller.DeleteTantangan)

	return r
}
