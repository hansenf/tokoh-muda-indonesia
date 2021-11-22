package routes

import (
	"net/http"
	"tmi/controller"
	"tmi/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	// middleware
	r.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
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

	//route leaderboard
	r.GET("/api/leaderboard", controller.GetAllLeaderboard)
	authorized.POST("/api/leaderboard", controller.CreateLeaderboard)
	r.GET("/api/leaderboard/:id", controller.GetLeaderboardId)
	authorized.PUT("/api/leaderboard/:id", controller.UpdateSilabus)
	authorized.DELETE("/api/leaderboard/:id", controller.DeleteSilabus)

	//route silabus
	r.GET("/api/silabus", controller.GetAllSilabus)
	authorized.POST("/api/silabus", controller.CreateSilabus)
	r.GET("/api/silabus/:id", controller.GetSilabusID)
	authorized.PUT("/api/silabus/:id", controller.UpdateLeaderboard)
	authorized.DELETE("/api/silabus/:id", controller.DeleteLeaderboard)

	return r
}
