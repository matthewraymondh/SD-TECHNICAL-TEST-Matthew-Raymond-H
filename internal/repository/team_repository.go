package repository

import (
	"xyz-football-api/internal/model"

	"gorm.io/gorm"
)

// TeamRepository menangani operasi database untuk Team
type TeamRepository struct {
	db *gorm.DB
}

// NewTeamRepository membuat instance TeamRepository baru
func NewTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{db: db}
}

// Create membuat team baru
func (r *TeamRepository) Create(team *model.Team) error {
	return r.db.Create(team).Error
}

// FindAll mengambil semua team
func (r *TeamRepository) FindAll() ([]model.Team, error) {
	var teams []model.Team
	err := r.db.Find(&teams).Error
	return teams, err
}

// FindByID mengambil team berdasarkan ID
func (r *TeamRepository) FindByID(id uint) (*model.Team, error) {
	var team model.Team
	err := r.db.First(&team, id).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}

// Update memperbarui data team
func (r *TeamRepository) Update(team *model.Team) error {
	return r.db.Save(team).Error
}

// Delete menghapus team (soft delete)
func (r *TeamRepository) Delete(id uint) error {
	return r.db.Delete(&model.Team{}, id).Error
}

// CountWinsByTeamID menghitung total kemenangan tim sebagai home atau away
func (r *TeamRepository) CountWinsByTeamID(teamID uint) (int64, error) {
	var count int64

	// Hitung kemenangan sebagai home team (home_score > away_score)
	err := r.db.Model(&model.Match{}).
		Where("home_team_id = ? AND status = ? AND home_score > away_score", teamID, model.MatchStatusCompleted).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	homeWins := count

	// Hitung kemenangan sebagai away team (away_score > home_score)
	err = r.db.Model(&model.Match{}).
		Where("away_team_id = ? AND status = ? AND away_score > home_score", teamID, model.MatchStatusCompleted).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	awayWins := count

	return homeWins + awayWins, nil
}
