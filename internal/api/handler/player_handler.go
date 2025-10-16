package handler

import (
	"net/http"
	"strconv"
	"xyz-football-api/internal/model"
	"xyz-football-api/internal/repository"
	"xyz-football-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

// PlayerHandler menangani endpoint players
type PlayerHandler struct {
	playerRepo *repository.PlayerRepository
	teamRepo   *repository.TeamRepository
}

// NewPlayerHandler membuat instance PlayerHandler baru
func NewPlayerHandler(playerRepo *repository.PlayerRepository, teamRepo *repository.TeamRepository) *PlayerHandler {
	return &PlayerHandler{
		playerRepo: playerRepo,
		teamRepo:   teamRepo,
	}
}

// CreatePlayerRequest adalah struct untuk request body create player
type CreatePlayerRequest struct {
	Name         string `json:"name" binding:"required"`
	TeamID       uint   `json:"team_id" binding:"required"`
	Position     string `json:"position" binding:"required,oneof=penyerang gelandang bertahan 'penjaga gawang'"`
	JerseyNumber int    `json:"jersey_number" binding:"required,min=1,max=99"`
	HeightCm     *int   `json:"height_cm"`
	WeightKg     *int   `json:"weight_kg"`
}

// CreatePlayer menangani endpoint POST /players
// @Summary Membuat player baru
// @Description Endpoint untuk membuat player baru dalam sebuah team
// @Tags Players
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body CreatePlayerRequest true "Player Data"
// @Success 201 {object} model.Player
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /players [post]
func (h *PlayerHandler) CreatePlayer(c *gin.Context) {
	var req CreatePlayerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Data player tidak valid: "+err.Error())
		return
	}

	// Mapping request ke model
	player := model.Player{
		Name:         req.Name,
		TeamID:       req.TeamID,
		Position:     req.Position,
		JerseyNumber: req.JerseyNumber,
		HeightCm:     req.HeightCm,
		WeightKg:     req.WeightKg,
	}

	// Validasi team exists
	_, err := h.teamRepo.FindByID(player.TeamID)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Team tidak ditemukan")
		return
	}

	// Cek duplikat jersey number
	exists, err := h.playerRepo.CheckJerseyNumberExists(player.TeamID, player.JerseyNumber, 0)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Gagal validasi nomor punggung: "+err.Error())
		return
	}
	if exists {
		utils.RespondError(c, http.StatusBadRequest, "Nomor punggung sudah digunakan di tim ini")
		return
	}

	if err := h.playerRepo.Create(&player); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Gagal membuat player: "+err.Error())
		return
	}

	utils.RespondSuccess(c, http.StatusCreated, player)
}

// GetPlayersByTeam menangani endpoint GET /teams/:id/players
// @Summary Mengambil semua players dari team tertentu
// @Description Endpoint untuk mengambil daftar semua player dalam sebuah team
// @Tags Players
// @Produce json
// @Security BearerAuth
// @Param id path int true "Team ID"
// @Success 200 {array} model.Player
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /teams/{id}/players [get]
func (h *PlayerHandler) GetPlayersByTeam(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "ID team tidak valid")
		return
	}

	// Validasi team exists
	_, err = h.teamRepo.FindByID(uint(id))
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, "Team tidak ditemukan")
		return
	}

	players, err := h.playerRepo.FindByTeamID(uint(id))
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Gagal mengambil data players: "+err.Error())
		return
	}

	utils.RespondSuccess(c, http.StatusOK, players)
}

// UpdatePlayer menangani endpoint PUT /players/:id
// @Summary Memperbarui data player
// @Description Endpoint untuk memperbarui informasi player
// @Tags Players
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Player ID"
// @Param body body model.Player true "Updated Player Data"
// @Success 200 {object} model.Player
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /players/{id} [put]
func (h *PlayerHandler) UpdatePlayer(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "ID player tidak valid")
		return
	}

	// Cek apakah player ada
	existingPlayer, err := h.playerRepo.FindByID(uint(id))
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, "Player tidak ditemukan")
		return
	}

	// Bind data baru (tidak perlu binding required untuk update partial)
	type UpdatePlayerRequest struct {
		Name         *string `json:"name,omitempty"`
		HeightCm     *int    `json:"height_cm,omitempty"`
		WeightKg     *int    `json:"weight_kg,omitempty"`
		Position     *string `json:"position,omitempty"`
		JerseyNumber *int    `json:"jersey_number,omitempty"`
	}

	var updateData UpdatePlayerRequest
	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Data player tidak valid: "+err.Error())
		return
	}

	// Update fields jika ada
	if updateData.Name != nil {
		existingPlayer.Name = *updateData.Name
	}
	if updateData.HeightCm != nil {
		existingPlayer.HeightCm = updateData.HeightCm
	}
	if updateData.WeightKg != nil {
		existingPlayer.WeightKg = updateData.WeightKg
	}
	if updateData.Position != nil {
		existingPlayer.Position = *updateData.Position
	}
	if updateData.JerseyNumber != nil {
		// Cek duplikat jersey number
		exists, err := h.playerRepo.CheckJerseyNumberExists(existingPlayer.TeamID, *updateData.JerseyNumber, existingPlayer.ID)
		if err != nil {
			utils.RespondError(c, http.StatusInternalServerError, "Gagal validasi nomor punggung: "+err.Error())
			return
		}
		if exists {
			utils.RespondError(c, http.StatusBadRequest, "Nomor punggung sudah digunakan di tim ini")
			return
		}
		existingPlayer.JerseyNumber = *updateData.JerseyNumber
	}

	if err := h.playerRepo.Update(existingPlayer); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Gagal memperbarui player: "+err.Error())
		return
	}

	utils.RespondSuccess(c, http.StatusOK, existingPlayer)
}

// DeletePlayer menangani endpoint DELETE /players/:id
// @Summary Menghapus player
// @Description Endpoint untuk menghapus player (soft delete)
// @Tags Players
// @Produce json
// @Security BearerAuth
// @Param id path int true "Player ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /players/{id} [delete]
func (h *PlayerHandler) DeletePlayer(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "ID player tidak valid")
		return
	}

	if err := h.playerRepo.Delete(uint(id)); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Gagal menghapus player: "+err.Error())
		return
	}

	utils.RespondMessage(c, http.StatusOK, "Player deleted successfully")
}
