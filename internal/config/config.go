package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	JWTSecret string
	DBPath    string
	Env       string // "development" atau "production"
}

var AppConfig *Config

func LoadConfig() *Config {
	// Coba load .env, tapi kalau tidak ada file-nya (misal di production
	// yang env-nya di-set langsung di server), jangan sampai app crash
	if err := godotenv.Load(); err != nil {
		log.Println("Tidak ada file .env ditemukan, menggunakan environment variable sistem")
	}

	cfg := &Config{
		Port:      getEnv("PORT", "8080"),
		JWTSecret: getEnv("JWT_SECRET", ""),
		DBPath:    getEnv("DB_PATH", "portfolio.db"),
		Env:       getEnv("APP_ENV", "development"),
	}

	// Validasi: kalau JWT_SECRET kosong, hentikan app sekarang juga.
	// Lebih baik crash saat startup daripada jalan dengan auth yang rapuh.
	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET wajib diisi di file .env")
	}

	AppConfig = cfg
	return cfg
}

// getEnv ambil env variable, kalau kosong pakai nilai default
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}