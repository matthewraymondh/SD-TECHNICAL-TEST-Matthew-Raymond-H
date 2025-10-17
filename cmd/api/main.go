package main

import (
	"fmt"
	"log"
	"xyz-football-api/config"
	"xyz-football-api/internal/api"
	"xyz-football-api/pkg/database"
)

func main() {
	// ASCII Art Banner
	banner := `
╔═══════════════════════════════════════════════════════════╗
║                                                           ║
║        XYZ FOOTBALL API - Backend Management System       ║
║                                                           ║
║        Version: 1.0.0                                     ║
║        Author: Matthew Raymond Hartono                   	║
║                                                           ║
╚═══════════════════════════════════════════════════════════╝
	`
	fmt.Println(banner)

	// Load configuration
	log.Println("⏳ Loading configuration...")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}
	log.Println("✓ Configuration loaded successfully")

	// Initialize database connection
	log.Println("⏳ Connecting to database...")
	if err := database.InitDatabase(&cfg.Database); err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	// Run auto migration
	log.Println("⏳ Running database migrations...")
	if err := database.AutoMigrate(); err != nil {
		log.Fatalf("❌ Failed to run migrations: %v", err)
	}

	// Setup router
	log.Println("⏳ Setting up routes...")
	router := api.SetupRouter(cfg, database.GetDB())
	log.Println("✓ Routes configured successfully")

	// Start server
	serverAddress := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("\n🚀 Server is running on http://localhost%s\n", serverAddress)
	log.Println("📝 Press CTRL+C to stop the server")
	log.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	if err := router.Run(serverAddress); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
