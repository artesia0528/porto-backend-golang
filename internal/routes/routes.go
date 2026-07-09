package routes

import (
	"portfolio-backend/internal/handlers"
	"portfolio-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes mendaftarkan semua route ke Gin engine.
func SetupRoutes(r *gin.Engine, jwtSecret string, authHandler *handlers.AuthHandler, projectHandler *handlers.ProjectHandler) {
	api := r.Group("/api")
	{
		// Publik — tanpa login
		api.POST("/login", authHandler.Login)
		api.POST("/register", authHandler.Register)
		api.GET("/projects", projectHandler.GetProjects)

		// Butuh login (untuk content management)
		admin := api.Group("/admin")
		admin.Use(middleware.AuthRequired(jwtSecret))
		{
			admin.POST("/projects", projectHandler.CreateProject)
			admin.PUT("/projects/:id", projectHandler.UpdateProject)
			admin.DELETE("/projects/:id", projectHandler.DeleteProject)
		}
	}
}