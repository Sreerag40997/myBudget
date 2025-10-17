package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sreerag/myBudget/controllers"
	"github.com/sreerag/myBudget/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		protected := api.Group("/")
		protected.Use(middleware.AuthRequired())
		{
			protected.GET("/profile", func(c *gin.Context) {
				uid := c.GetUint("user_id")
				email := c.GetString("email")
				c.JSON(200, gin.H{"user_id": uid, "email": email})
			})
		}
	}

	return r
}
