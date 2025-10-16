package middleware

import (
	"net/http"
	"strings"
	"xyz-football-api/config"
	"xyz-football-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware memvalidasi JWT token dari header Authorization
func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			utils.RespondError(c, http.StatusUnauthorized, "Token tidak ditemukan")
			c.Abort()
			return
		}

		// Format: "Bearer <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.RespondError(c, http.StatusUnauthorized, "Format token tidak valid")
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Validasi token
		claims, err := utils.ValidateJWT(tokenString, cfg.JWT.Secret)
		if err != nil {
			utils.RespondError(c, http.StatusUnauthorized, "Token tidak valid atau sudah expired")
			c.Abort()
			return
		}

		// Simpan user info ke context untuk digunakan handler
		c.Set("username", claims.Username)

		c.Next()
	}
}
