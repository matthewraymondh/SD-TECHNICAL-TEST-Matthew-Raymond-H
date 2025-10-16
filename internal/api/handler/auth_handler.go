package handler

import (
	"net/http"
	"xyz-football-api/config"
	"xyz-football-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

// AuthHandler menangani endpoint autentikasi
type AuthHandler struct {
	cfg *config.Config
}

// NewAuthHandler membuat instance AuthHandler baru
func NewAuthHandler(cfg *config.Config) *AuthHandler {
	return &AuthHandler{cfg: cfg}
}

// LoginRequest merepresentasikan struktur request login
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse merepresentasikan struktur response login
type LoginResponse struct {
	Token string `json:"token"`
}

// Login menangani endpoint POST /login
// @Summary Login untuk mendapatkan JWT token
// @Description Endpoint untuk autentikasi menggunakan username dan password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param body body LoginRequest true "Credentials"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Data request tidak valid")
		return
	}

	// Validasi credentials dengan data dari config
	if req.Username != h.cfg.Admin.Username || req.Password != h.cfg.Admin.Password {
		utils.RespondError(c, http.StatusUnauthorized, "Username atau password salah")
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(
		req.Username,
		h.cfg.JWT.Secret,
		h.cfg.JWT.ExpirationHours,
	)

	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Gagal membuat token")
		return
	}

	utils.RespondSuccess(c, http.StatusOK, LoginResponse{Token: token})
}
