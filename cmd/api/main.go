package main

import (
	"portfolio-backend/internal/config"
	"portfolio-backend/internal/database"
	"portfolio-backend/internal/middleware"
	"portfolio-backend/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig() // load & validasi config paling pertama

	database.Connect()

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	routes.SetupRoutes(r)

	r.Run(":" + cfg.Port)
}