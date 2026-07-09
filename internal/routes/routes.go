package routes

import (
	"portfolio-backend/internal/handlers"
	"portfolio-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		// Publik — tanpa login
		api.GET("/projects", handlers.GetProjects)
		api.POST("/login", handlers.Login)

		// Butuh login (untuk content management)
		admin := api.Group("/admin")
		admin.Use(middleware.AuthRequired())
		{
			admin.POST("/projects", handlers.CreateProject)
			admin.DELETE("/projects/:id", handlers.DeleteProject)
		}
	}
}