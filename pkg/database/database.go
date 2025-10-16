package database

import (
	"fmt"
	"log"
	"xyz-football-api/config"
	"xyz-football-api/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB adalah instance global database
var DB *gorm.DB

// InitDatabase menginisialisasi koneksi ke database PostgreSQL
func InitDatabase(cfg *config.DatabaseConfig) error {
	dsn := cfg.GetDSN()

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return fmt.Errorf("gagal koneksi ke database: %w", err)
	}

	log.Println("✓ Koneksi database berhasil")

	return nil
}

// AutoMigrate menjalankan migrasi otomatis untuk semua model
func AutoMigrate() error {
	log.Println("Menjalankan auto migration...")

	err := DB.AutoMigrate(
		&model.Team{},
		&model.Player{},
		&model.Match{},
		&model.Goal{},
	)

	if err != nil {
		return fmt.Errorf("gagal melakukan auto migrate: %w", err)
	}

	log.Println("✓ Auto migration selesai")
	return nil
}

// GetDB mengembalikan instance database
func GetDB() *gorm.DB {
	return DB
}
