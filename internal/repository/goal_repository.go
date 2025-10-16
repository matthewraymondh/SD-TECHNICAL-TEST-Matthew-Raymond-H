package repository

import (
	"xyz-football-api/internal/model"

	"gorm.io/gorm"
)

// GoalRepository menangani operasi database untuk Goal
type GoalRepository struct {
	db *gorm.DB
}

// NewGoalRepository membuat instance GoalRepository baru
func NewGoalRepository(db *gorm.DB) *GoalRepository {
	return &GoalRepository{db: db}
}

// CreateBatch membuat multiple goals sekaligus
func (r *GoalRepository) CreateBatch(goals []model.Goal) error {
	if len(goals) == 0 {
		return nil
	}
	return r.db.Create(&goals).Error
}

// FindByMatchID mengambil semua goal dari match tertentu
func (r *GoalRepository) FindByMatchID(matchID uint) ([]model.Goal, error) {
	var goals []model.Goal
	err := r.db.Where("match_id = ?", matchID).Preload("Player").Find(&goals).Error
	return goals, err
}

// GetTopScorerInMatch mengambil pencetak gol terbanyak dalam satu pertandingan
func (r *GoalRepository) GetTopScorerInMatch(matchID uint) (*model.Player, int, error) {
	type Result struct {
		PlayerID  uint
		GoalCount int
	}

	var result Result
	err := r.db.Model(&model.Goal{}).
		Select("player_id, COUNT(*) as goal_count").
		Where("match_id = ?", matchID).
		Group("player_id").
		Order("goal_count DESC").
		Limit(1).
		Scan(&result).Error

	if err != nil {
		return nil, 0, err
	}

	// Jika tidak ada gol
	if result.PlayerID == 0 {
		return nil, 0, nil
	}

	// Ambil data player
	var player model.Player
	err = r.db.First(&player, result.PlayerID).Error
	if err != nil {
		return nil, 0, err
	}

	return &player, result.GoalCount, nil
}

// DeleteByMatchID menghapus semua goal dari match tertentu (untuk update result)
func (r *GoalRepository) DeleteByMatchID(matchID uint) error {
	return r.db.Where("match_id = ?", matchID).Delete(&model.Goal{}).Error
}
