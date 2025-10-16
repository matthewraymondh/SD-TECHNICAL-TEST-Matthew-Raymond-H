package repository

import (
	"xyz-football-api/internal/model"

	"gorm.io/gorm"
)

// PlayerRepository menangani operasi database untuk Player
type PlayerRepository struct {
	db *gorm.DB
}

// NewPlayerRepository membuat instance PlayerRepository baru
func NewPlayerRepository(db *gorm.DB) *PlayerRepository {
	return &PlayerRepository{db: db}
}

// Create membuat player baru
func (r *PlayerRepository) Create(player *model.Player) error {
	return r.db.Create(player).Error
}

// FindByTeamID mengambil semua player dari team tertentu
func (r *PlayerRepository) FindByTeamID(teamID uint) ([]model.Player, error) {
	var players []model.Player
	err := r.db.Where("team_id = ?", teamID).Find(&players).Error
	return players, err
}

// FindByID mengambil player berdasarkan ID
func (r *PlayerRepository) FindByID(id uint) (*model.Player, error) {
	var player model.Player
	err := r.db.First(&player, id).Error
	if err != nil {
		return nil, err
	}
	return &player, nil
}

// Update memperbarui data player
func (r *PlayerRepository) Update(player *model.Player) error {
	return r.db.Save(player).Error
}

// Delete menghapus player (soft delete)
func (r *PlayerRepository) Delete(id uint) error {
	return r.db.Delete(&model.Player{}, id).Error
}

// CheckJerseyNumberExists memeriksa apakah nomor punggung sudah digunakan di tim
func (r *PlayerRepository) CheckJerseyNumberExists(teamID uint, jerseyNumber int, excludePlayerID uint) (bool, error) {
	var count int64
	query := r.db.Model(&model.Player{}).
		Where("team_id = ? AND jersey_number = ?", teamID, jerseyNumber)

	// Exclude player tertentu jika sedang update
	if excludePlayerID > 0 {
		query = query.Where("id != ?", excludePlayerID)
	}

	err := query.Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
