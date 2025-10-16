package model

import (
	"time"
)

// Goal merepresentasikan tabel goals di database
type Goal struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	MatchID   uint      `gorm:"not null" json:"match_id" binding:"required"`
	PlayerID  uint      `gorm:"not null" json:"player_id" binding:"required"`
	GoalTime  int       `gorm:"not null" json:"goal_time" binding:"required,min=1,max=120"`
	CreatedAt time.Time `json:"created_at"`

	// Relasi
	Match  Match  `gorm:"foreignKey:MatchID" json:"-"`
	Player Player `gorm:"foreignKey:PlayerID" json:"player,omitempty"`
}

// TableName menentukan nama tabel untuk model Goal
func (Goal) TableName() string {
	return "goals"
}
