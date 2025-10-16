package api

import (
	"xyz-football-api/config"
	"xyz-football-api/internal/api/handler"
	"xyz-football-api/internal/api/middleware"
	"xyz-football-api/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRouter mengkonfigurasi semua routes dan middleware
func SetupRouter(cfg *config.Config, db *gorm.DB) *gin.Engine {
	// Set Gin mode
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	// Apply global middleware
	router.Use(gin.Recovery())
	router.Use(middleware.Logger())
	router.Use(middleware.CORS())

	// Initialize repositories
	teamRepo := repository.NewTeamRepository(db)
	playerRepo := repository.NewPlayerRepository(db)
	matchRepo := repository.NewMatchRepository(db)
	goalRepo := repository.NewGoalRepository(db)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(cfg)
	teamHandler := handler.NewTeamHandler(teamRepo)
	playerHandler := handler.NewPlayerHandler(playerRepo, teamRepo)
	matchHandler := handler.NewMatchHandler(matchRepo, teamRepo, playerRepo, goalRepo)

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "XYZ Football API is running",
		})
	})

	// Public routes (tidak memerlukan autentikasi)
	router.POST("/login", authHandler.Login)

	// Protected routes (memerlukan JWT token)
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware(cfg))
	{
		// Teams endpoints
		protected.POST("/teams", teamHandler.CreateTeam)
		protected.GET("/teams", teamHandler.GetAllTeams)
		protected.GET("/teams/:id", teamHandler.GetTeamByID)
		protected.PUT("/teams/:id", teamHandler.UpdateTeam)
		protected.DELETE("/teams/:id", teamHandler.DeleteTeam)

		// Players endpoints
		protected.POST("/players", playerHandler.CreatePlayer)
		protected.GET("/teams/:id/players", playerHandler.GetPlayersByTeam)
		protected.PUT("/players/:id", playerHandler.UpdatePlayer)
		protected.DELETE("/players/:id", playerHandler.DeletePlayer)

		// Matches endpoints
		protected.POST("/matches", matchHandler.CreateMatch)
		protected.POST("/matches/:id/result", matchHandler.ReportMatchResult)
		protected.GET("/matches/:id/report", matchHandler.GetMatchReport)
	}

	return router
}
