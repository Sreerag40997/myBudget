package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sreerag/myBudget/controllers"
	"github.com/sreerag/myBudget/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/admin/login", func(c *gin.Context) {
		c.File("./templates/admin-login.html")
	})
	r.GET("/admin/dashboard", func(c *gin.Context) {
		c.File("./templates/admin/dashboard.html")
	})

	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
		api.GET("/profile", middleware.AuthRequired(), controllers.GetProfile)
		api.PUT("/profile", middleware.AuthRequired(), controllers.UpdateProfile)

		api.POST("/admin/login", controllers.AdminLogin)

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
