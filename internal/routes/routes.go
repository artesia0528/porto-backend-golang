package routes

import (
	"portfolio-backend/internal/handlers"
	"portfolio-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes mendaftarkan semua route ke Gin engine.
func SetupRoutes(r *gin.Engine, jwtSecret string, authHandler *handlers.AuthHandler, projectHandler *handlers.ProjectHandler, messageHandler *handlers.MessageHandler, blogHandler *handlers.BlogHandler) {
	api := r.Group("/api")
	{
		// Publik — tanpa login
		api.POST("/login", authHandler.Login)
		// api.POST("/register", authHandler.Register)
		api.GET("/projects", projectHandler.GetProjects)
		
		api.POST("/contact", messageHandler.Create)

		api.GET("/blogs", blogHandler.GetAll)

		// Butuh login (untuk content management)
		admin := api.Group("/admin")
		admin.Use(middleware.AuthRequired(jwtSecret))
		{
			// Projects
			admin.POST("/projects", projectHandler.CreateProject)
			admin.PUT("/projects/:id", projectHandler.UpdateProject)
			admin.DELETE("/projects/:id", projectHandler.DeleteProject)

			// Messages
			admin.GET("/messages", messageHandler.GetAll)
			admin.PATCH("/messages/:id/read", messageHandler.MarkAsRead)
			admin.DELETE("/messages/:id", messageHandler.Delete)

			//BLog
			admin.POST("/blogs", blogHandler.Create)
            admin.PUT("/blogs/:id", blogHandler.Update)
            admin.DELETE("/blogs/:id", blogHandler.Delete)
		}
	}
}