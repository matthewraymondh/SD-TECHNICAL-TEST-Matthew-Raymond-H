package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"xyz-football-api/internal/model"
	"xyz-football-api/internal/repository"
	"xyz-football-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

// MatchHandler menangani endpoint matches
type MatchHandler struct {
	matchRepo  *repository.MatchRepository
	teamRepo   *repository.TeamRepository
	playerRepo *repository.PlayerRepository
	goalRepo   *repository.GoalRepository
}

// NewMatchHandler membuat instance MatchHandler baru
func NewMatchHandler(
	matchRepo *repository.MatchRepository,
	teamRepo *repository.TeamRepository,
	playerRepo *repository.PlayerRepository,
	goalRepo *repository.GoalRepository,
) *MatchHandler {
	return &MatchHandler{
		matchRepo:  matchRepo,
		teamRepo:   teamRepo,
		playerRepo: playerRepo,
		goalRepo:   goalRepo,
	}
}

// CreateMatchRequest adalah struct untuk request body create match
type CreateMatchRequest struct {
	HomeTeamID    uint   `json:"home_team_id" binding:"required"`
	AwayTeamID    uint   `json:"away_team_id" binding:"required"`
	MatchDatetime string `json:"match_datetime" binding:"required"`
}

// CreateMatch menangani endpoint POST /matches
// @Summary Membuat jadwal pertandingan baru
// @Description Endpoint untuk membuat jadwal pertandingan baru
// @Tags Matches
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body CreateMatchRequest true "Match Data"
// @Success 201 {object} model.Match
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /matches [post]
func (h *MatchHandler) CreateMatch(c *gin.Context) {
	var req CreateMatchRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Data match tidak valid: "+err.Error())
		return
	}

	// Parse match datetime
	matchTime, err := time.Parse(time.RFC3339, req.MatchDatetime)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Format match_datetime tidak valid (gunakan ISO 8601/RFC3339)")
		return
	}

	// Mapping request ke model
	match := model.Match{
		HomeTeamID:    req.HomeTeamID,
		AwayTeamID:    req.AwayTeamID,
		MatchDatetime: matchTime,
	}

	// Validasi home team exists
	_, err = h.teamRepo.FindByID(match.HomeTeamID)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Home team tidak ditemukan")
		return
	}

	// Validasi away team exists
	_, err = h.teamRepo.FindByID(match.AwayTeamID)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Away team tidak ditemukan")
		return
	}

	// Validasi home team != away team
	if match.HomeTeamID == match.AwayTeamID {
		utils.RespondError(c, http.StatusBadRequest, "Home team dan away team tidak boleh sama")
		return
	}

	// Set default status
	match.Status = model.MatchStatusScheduled
	match.HomeScore = 0
	match.AwayScore = 0

	if err := h.matchRepo.Create(&match); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Gagal membuat match: "+err.Error())
		return
	}

	utils.RespondSuccess(c, http.StatusCreated, match)
}

// ReportMatchResultRequest merepresentasikan request untuk melaporkan hasil pertandingan
type ReportMatchResultRequest struct {
	HomeScore int `json:"home_score" binding:"required,min=0"`
	AwayScore int `json:"away_score" binding:"required,min=0"`
	Goals     []struct {
		PlayerID uint `json:"player_id" binding:"required"`
		GoalTime int  `json:"goal_time" binding:"required,min=1,max=120"`
	} `json:"goals" binding:"required"`
}

// ReportMatchResult menangani endpoint POST /matches/:id/result
// @Summary Melaporkan hasil pertandingan
// @Description Endpoint untuk melaporkan hasil pertandingan dan mencatat gol
// @Tags Matches
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Match ID"
// @Param body body ReportMatchResultRequest true "Match Result Data"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /matches/{id}/result [post]
func (h *MatchHandler) ReportMatchResult(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "ID match tidak valid")
		return
	}

	// Validasi match exists
	match, err := h.matchRepo.FindByID(uint(id))
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, "Match tidak ditemukan")
		return
	}

	// Validasi status match
	if match.Status == model.MatchStatusCompleted {
		utils.RespondError(c, http.StatusBadRequest, "Match sudah dilaporkan sebelumnya")
		return
	}

	var req ReportMatchResultRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Data hasil match tidak valid: "+err.Error())
		return
	}

	// Validasi total gol
	totalGoals := req.HomeScore + req.AwayScore
	if len(req.Goals) != totalGoals {
		utils.RespondError(c, http.StatusBadRequest, fmt.Sprintf("Jumlah gol tidak sesuai. Total skor: %d, jumlah detail gol: %d", totalGoals, len(req.Goals)))
		return
	}

	// Validasi setiap player dalam goals
	for _, g := range req.Goals {
		player, err := h.playerRepo.FindByID(g.PlayerID)
		if err != nil {
			utils.RespondError(c, http.StatusBadRequest, fmt.Sprintf("Player dengan ID %d tidak ditemukan", g.PlayerID))
			return
		}

		// Validasi player adalah bagian dari salah satu tim yang bertanding
		if player.TeamID != match.HomeTeamID && player.TeamID != match.AwayTeamID {
			utils.RespondError(c, http.StatusBadRequest, fmt.Sprintf("Player %s tidak termasuk dalam tim yang bertanding", player.Name))
			return
		}
	}

	// Update match result
	if err := h.matchRepo.UpdateResult(uint(id), req.HomeScore, req.AwayScore); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Gagal memperbarui hasil match: "+err.Error())
		return
	}

	// Hapus goals lama jika ada (untuk update)
	h.goalRepo.DeleteByMatchID(uint(id))

	// Insert goals
	var goals []model.Goal
	for _, g := range req.Goals {
		goals = append(goals, model.Goal{
			MatchID:  uint(id),
			PlayerID: g.PlayerID,
			GoalTime: g.GoalTime,
		})
	}

	if err := h.goalRepo.CreateBatch(goals); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Gagal mencatat gol: "+err.Error())
		return
	}

	utils.RespondMessage(c, http.StatusOK, "Match result reported successfully")
}

// MatchReportResponse merepresentasikan response laporan pertandingan
type MatchReportResponse struct {
	Schedule          string `json:"schedule"`
	HomeTeam          string `json:"home_team"`
	AwayTeam          string `json:"away_team"`
	FinalScore        string `json:"final_score"`
	MatchResult       string `json:"match_result"`
	TopScorerInMatch  string `json:"top_scorer_in_match"`
	HomeTeamTotalWins int64  `json:"home_team_total_wins"`
	AwayTeamTotalWins int64  `json:"away_team_total_wins"`
}

// GetMatchReport menangani endpoint GET /matches/:id/report
// @Summary Mengambil laporan pertandingan
// @Description Endpoint untuk mengambil laporan detail pertandingan
// @Tags Matches
// @Produce json
// @Security BearerAuth
// @Param id path int true "Match ID"
// @Success 200 {object} MatchReportResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /matches/{id}/report [get]
func (h *MatchHandler) GetMatchReport(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "ID match tidak valid")
		return
	}

	// Ambil match dengan relasi
	match, err := h.matchRepo.FindByIDWithGoals(uint(id))
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, "Match tidak ditemukan")
		return
	}

	// Build response
	response := MatchReportResponse{
		Schedule:   match.MatchDatetime.Format("2006-01-02T15:04:05Z07:00"),
		HomeTeam:   match.HomeTeam.Name,
		AwayTeam:   match.AwayTeam.Name,
		FinalScore: fmt.Sprintf("%d-%d", match.HomeScore, match.AwayScore),
	}

	// Determine match result
	if match.Status != model.MatchStatusCompleted {
		response.MatchResult = "Belum Selesai"
	} else {
		if match.HomeScore > match.AwayScore {
			response.MatchResult = "Tim Home Menang"
		} else if match.AwayScore > match.HomeScore {
			response.MatchResult = "Tim Away Menang"
		} else {
			response.MatchResult = "Seri"
		}
	}

	// Get top scorer in match
	topScorer, goalCount, err := h.goalRepo.GetTopScorerInMatch(uint(id))
	if err == nil && topScorer != nil {
		response.TopScorerInMatch = fmt.Sprintf("%s (%d gol)", topScorer.Name, goalCount)
	} else {
		response.TopScorerInMatch = "Belum ada gol"
	}

	// Get total wins for each team
	homeWins, err := h.teamRepo.CountWinsByTeamID(match.HomeTeamID)
	if err == nil {
		response.HomeTeamTotalWins = homeWins
	}

	awayWins, err := h.teamRepo.CountWinsByTeamID(match.AwayTeamID)
	if err == nil {
		response.AwayTeamTotalWins = awayWins
	}

	utils.RespondSuccess(c, http.StatusOK, response)
}
