package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sreerag/myBudget/controllers"
	"github.com/sreerag/myBudget/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Serve admin login page (GET)
	r.GET("/admin/login", func(c *gin.Context) {
		c.File("./templates/admin-login.html")
	})
	r.GET("/admin/dashboard", func(c *gin.Context) {
		c.File("./templates/admin/dashboard.html")
	})

	// API routes
	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)      // /api/register
		api.POST("/login", controllers.Login)            // /api/login
		api.POST("/admin/login", controllers.AdminLogin) // /api/admin/login

		api.GET("/profile", middleware.AuthRequired(), controllers.GetProfile)
		api.PUT("/profile", middleware.AuthRequired(), controllers.UpdateProfile)

		admin := api.Group("/admin")
		admin.Use(middleware.AdminRequired())
		{
			admin.GET("/dashboard", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Welcome Admin"})
			})
		}
	}

	return r
}
