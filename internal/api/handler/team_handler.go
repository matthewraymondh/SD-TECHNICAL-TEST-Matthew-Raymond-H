package handler

import (
	"net/http"
	"strconv"
	"xyz-football-api/internal/model"
	"xyz-football-api/internal/repository"
	"xyz-football-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

// TeamHandler menangani endpoint teams
type TeamHandler struct {
	teamRepo *repository.TeamRepository
}

// NewTeamHandler membuat instance TeamHandler baru
func NewTeamHandler(teamRepo *repository.TeamRepository) *TeamHandler {
	return &TeamHandler{teamRepo: teamRepo}
}

// CreateTeam menangani endpoint POST /teams
// @Summary Membuat team baru
// @Description Endpoint untuk membuat team sepak bola baru
// @Tags Teams
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body model.Team true "Team Data"
// @Success 201 {object} model.Team
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /teams [post]
func (h *TeamHandler) CreateTeam(c *gin.Context) {
	var team model.Team

	if err := c.ShouldBindJSON(&team); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Data team tidak valid: "+err.Error())
		return
	}

	if err := h.teamRepo.Create(&team); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Gagal membuat team: "+err.Error())
		return
	}

	utils.RespondSuccess(c, http.StatusCreated, team)
}

// GetAllTeams menangani endpoint GET /teams
// @Summary Mengambil semua teams
// @Description Endpoint untuk mengambil daftar semua team
// @Tags Teams
// @Produce json
// @Security BearerAuth
// @Success 200 {array} model.Team
// @Failure 500 {object} utils.ErrorResponse
// @Router /teams [get]
func (h *TeamHandler) GetAllTeams(c *gin.Context) {
	teams, err := h.teamRepo.FindAll()
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Gagal mengambil data teams: "+err.Error())
		return
	}

	utils.RespondSuccess(c, http.StatusOK, teams)
}

// GetTeamByID menangani endpoint GET /teams/:id
// @Summary Mengambil team berdasarkan ID
// @Description Endpoint untuk mengambil detail team berdasarkan ID
// @Tags Teams
// @Produce json
// @Security BearerAuth
// @Param id path int true "Team ID"
// @Success 200 {object} model.Team
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /teams/{id} [get]
func (h *TeamHandler) GetTeamByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "ID team tidak valid")
		return
	}

	team, err := h.teamRepo.FindByID(uint(id))
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, "Team tidak ditemukan")
		return
	}

	utils.RespondSuccess(c, http.StatusOK, team)
}

// UpdateTeam menangani endpoint PUT /teams/:id
// @Summary Memperbarui data team
// @Description Endpoint untuk memperbarui informasi team
// @Tags Teams
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Team ID"
// @Param body body model.Team true "Updated Team Data"
// @Success 200 {object} model.Team
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /teams/{id} [put]
func (h *TeamHandler) UpdateTeam(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "ID team tidak valid")
		return
	}

	// Cek apakah team ada
	existingTeam, err := h.teamRepo.FindByID(uint(id))
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, "Team tidak ditemukan")
		return
	}

	// Bind data baru
	var updateData model.Team
	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Data team tidak valid: "+err.Error())
		return
	}

	// Update fields
	existingTeam.Name = updateData.Name
	if updateData.LogoURL != nil {
		existingTeam.LogoURL = updateData.LogoURL
	}
	if updateData.FoundedYear != nil {
		existingTeam.FoundedYear = updateData.FoundedYear
	}
	if updateData.HeadquartersAddress != nil {
		existingTeam.HeadquartersAddress = updateData.HeadquartersAddress
	}
	if updateData.HeadquartersCity != nil {
		existingTeam.HeadquartersCity = updateData.HeadquartersCity
	}

	if err := h.teamRepo.Update(existingTeam); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Gagal memperbarui team: "+err.Error())
		return
	}

	utils.RespondSuccess(c, http.StatusOK, existingTeam)
}

// DeleteTeam menangani endpoint DELETE /teams/:id
// @Summary Menghapus team
// @Description Endpoint untuk menghapus team (soft delete)
// @Tags Teams
// @Produce json
// @Security BearerAuth
// @Param id path int true "Team ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /teams/{id} [delete]
func (h *TeamHandler) DeleteTeam(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "ID team tidak valid")
		return
	}

	if err := h.teamRepo.Delete(uint(id)); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Gagal menghapus team: "+err.Error())
		return
	}

	utils.RespondMessage(c, http.StatusOK, "Team deleted successfully")
}
