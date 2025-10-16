package model

import (
	"time"

	"gorm.io/gorm"
)

// MatchStatus merepresentasikan status pertandingan
type MatchStatus string

const (
	MatchStatusScheduled MatchStatus = "scheduled"
	MatchStatusCompleted MatchStatus = "completed"
	MatchStatusCancelled MatchStatus = "cancelled"
)

// Match merepresentasikan tabel matches di database
type Match struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	HomeTeamID    uint           `gorm:"not null" json:"home_team_id" binding:"required"`
	AwayTeamID    uint           `gorm:"not null" json:"away_team_id" binding:"required"`
	MatchDatetime time.Time      `gorm:"not null" json:"match_datetime" binding:"required"`
	Status        MatchStatus    `gorm:"type:varchar(50);not null;default:'scheduled';check:status IN ('scheduled', 'completed', 'cancelled')" json:"status"`
	HomeScore     int            `gorm:"default:0" json:"home_score"`
	AwayScore     int            `gorm:"default:0" json:"away_score"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	// Relasi
	HomeTeam Team   `gorm:"foreignKey:HomeTeamID" json:"home_team,omitempty"`
	AwayTeam Team   `gorm:"foreignKey:AwayTeamID" json:"away_team,omitempty"`
	Goals    []Goal `gorm:"foreignKey:MatchID" json:"goals,omitempty"`
}

// TableName menentukan nama tabel untuk model Match
func (Match) TableName() string {
	return "matches"
}
