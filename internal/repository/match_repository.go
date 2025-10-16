package repository

import (
	"xyz-football-api/internal/model"

	"gorm.io/gorm"
)

// MatchRepository menangani operasi database untuk Match
type MatchRepository struct {
	db *gorm.DB
}

// NewMatchRepository membuat instance MatchRepository baru
func NewMatchRepository(db *gorm.DB) *MatchRepository {
	return &MatchRepository{db: db}
}

// Create membuat match baru
func (r *MatchRepository) Create(match *model.Match) error {
	return r.db.Create(match).Error
}

// FindByID mengambil match berdasarkan ID dengan relasi teams
func (r *MatchRepository) FindByID(id uint) (*model.Match, error) {
	var match model.Match
	err := r.db.Preload("HomeTeam").Preload("AwayTeam").First(&match, id).Error
	if err != nil {
		return nil, err
	}
	return &match, nil
}

// FindByIDWithGoals mengambil match beserta goals dan player info
func (r *MatchRepository) FindByIDWithGoals(id uint) (*model.Match, error) {
	var match model.Match
	err := r.db.Preload("HomeTeam").
		Preload("AwayTeam").
		Preload("Goals.Player").
		First(&match, id).Error
	if err != nil {
		return nil, err
	}
	return &match, nil
}

// Update memperbarui data match
func (r *MatchRepository) Update(match *model.Match) error {
	return r.db.Save(match).Error
}

// UpdateResult memperbarui hasil pertandingan dalam transaksi
func (r *MatchRepository) UpdateResult(matchID uint, homeScore, awayScore int) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Update match scores dan status
		err := tx.Model(&model.Match{}).
			Where("id = ?", matchID).
			Updates(map[string]interface{}{
				"home_score": homeScore,
				"away_score": awayScore,
				"status":     model.MatchStatusCompleted,
			}).Error

		return err
	})
}
