package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"portfolio-backend/internal/config"
	"portfolio-backend/internal/database"
	"portfolio-backend/internal/handlers"
	"portfolio-backend/internal/middleware"
	"portfolio-backend/internal/repositories"
	"portfolio-backend/internal/routes"
	"portfolio-backend/internal/services"
	"portfolio-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Load config
	cfg := config.LoadConfig()

	slog.Info("Memulai server",
		"port", cfg.Port,
		"env", cfg.Env,
	)

	// 2. Koneksi database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Gagal konek database: ", err)
	}
	slog.Info("Database PostgreSQL terkoneksi")

	// 3. Seed admin user (jika belum ada)
	database.SeedAdmin(db)

	// 4. Inisialisasi direktori upload
	uploadDirs := []string{"projects", "blogs", "experiences"}
	if err := utils.InitUploadDirs(uploadDirs); err != nil {
		log.Fatal("Gagal membuat direktori upload: ", err)
	}
	slog.Info("Direktori upload siap")

	// 5. Setup Gin
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	// Serve file upload secara publik
	r.Static("/uploads", "./uploads")

	// 6. Dependency Injection: wire semua layer
	userRepo := repositories.NewUserRepository(db)
	projectRepo := repositories.NewProjectRepository(db)
	messageRepo := repositories.NewMessageRepository(db)
	blogRepo := repositories.NewBlogRepository(db)
	experienceRepo := repositories.NewExperienceRepository(db)

	authService := services.NewAuthService(userRepo, cfg.JWTSecret)
	projectService := services.NewProjectService(projectRepo, cfg.BaseURL)
	messageService := services.NewMessageService(messageRepo)
	blogService := services.NewBlogService(blogRepo, cfg.BaseURL)
	experienceService := services.NewExperienceService(experienceRepo, cfg.BaseURL)

	authHandler := handlers.NewAuthHandler(authService)
	projectHandler := handlers.NewProjectHandler(projectService)
	messageHandler := handlers.NewMessageHandler(messageService)
	blogHandler := handlers.NewBlogHandler(blogService)
	experienceHandler := handlers.NewExperienceHandler(experienceService)

	// 7. Register routes
	routes.SetupRoutes(r, cfg.JWTSecret, authHandler, projectHandler, messageHandler, blogHandler, experienceHandler)

	// 8. Graceful shutdown
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	// Start server di goroutine
	go func() {
		slog.Info("Server berjalan", "addr", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %s\n", err)
		}
	}()

	// Tunggu signal interrupt (Ctrl+C) atau SIGTERM
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("Mematikan server...")

	// Beri waktu 5 detik untuk menyelesaikan request yang sedang berjalan
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server dipaksa berhenti: ", err)
	}

	slog.Info("Server berhenti dengan baik")
}
