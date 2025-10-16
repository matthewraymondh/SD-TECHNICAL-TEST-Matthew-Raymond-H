package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config menyimpan semua konfigurasi aplikasi
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Admin    AdminConfig
}

// ServerConfig berisi konfigurasi server
type ServerConfig struct {
	Port string
}

// DatabaseConfig berisi konfigurasi database PostgreSQL
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	Timezone string
}

// JWTConfig berisi konfigurasi JSON Web Token
type JWTConfig struct {
	Secret          string
	ExpirationHours int
}

// AdminConfig berisi kredensial admin untuk testing
type AdminConfig struct {
	Username string
	Password string
}

// LoadConfig membaca file .env dan mengembalikan struktur Config
func LoadConfig() (*Config, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Parse JWT expiration hours
	jwtExpHours, err := strconv.Atoi(getEnv("JWT_EXPIRATION_HOURS", "24"))
	if err != nil {
		jwtExpHours = 24
	}

	config := &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "football_db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
			Timezone: getEnv("DB_TIMEZONE", "Asia/Jakarta"),
		},
		JWT: JWTConfig{
			Secret:          getEnv("JWT_SECRET", "default_secret_change_this"),
			ExpirationHours: jwtExpHours,
		},
		Admin: AdminConfig{
			Username: getEnv("ADMIN_USERNAME", "admin"),
			Password: getEnv("ADMIN_PASSWORD", "admin123"),
		},
	}

	// Validasi konfigurasi penting
	if config.Database.Password == "" {
		return nil, fmt.Errorf("DB_PASSWORD tidak boleh kosong")
	}

	if config.JWT.Secret == "default_secret_change_this" {
		log.Println("WARNING: Menggunakan JWT_SECRET default. Ganti ini di production!")
	}

	return config, nil
}

// GetDSN mengembalikan Data Source Name untuk koneksi PostgreSQL
func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode, c.Timezone,
	)
}

// getEnv membaca environment variable atau mengembalikan nilai default
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
