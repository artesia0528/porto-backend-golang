package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config menyimpan semua konfigurasi aplikasi.
type Config struct {
	Port        string
	JWTSecret   string
	DatabaseURL string
	Env         string // "development" atau "production"
	BaseURL     string // URL dasar untuk serving file (misal: http://localhost:8080)
}

// LoadConfig membaca konfigurasi dari environment variables dan .env file.
// Mengembalikan *Config — caller bertanggung jawab menyimpan referensinya.
func LoadConfig() *Config {
	// Coba load .env, tapi kalau tidak ada file-nya (misal di production
	// yang env-nya di-set langsung di server), jangan sampai app crash
	if err := godotenv.Load(); err != nil {
		log.Println("Tidak ada file .env ditemukan, menggunakan environment variable sistem")
	}

	cfg := &Config{
		Port:        getEnv("PORT", "8080"),
		JWTSecret:   getEnv("JWT_SECRET", ""),
		DatabaseURL: getEnv("DATABASE_URL", ""),
		Env:         getEnv("APP_ENV", "development"),
		BaseURL:     getEnv("BASE_URL", "http://localhost:8080"),
	}

	// Validasi: kalau JWT_SECRET kosong, hentikan app sekarang juga.
	// Lebih baik crash saat startup daripada jalan dengan auth yang rapuh.
	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET wajib diisi di file .env")
	}

	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL wajib diisi di file .env")
	}

	return cfg
}

// getEnv ambil env variable, kalau kosong pakai nilai default.
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}