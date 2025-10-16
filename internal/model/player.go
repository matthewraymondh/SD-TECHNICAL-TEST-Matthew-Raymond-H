package model

import (
	"time"

	"gorm.io/gorm"
)

// Player merepresentasikan tabel players di database
type Player struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	TeamID       uint           `gorm:"not null" json:"team_id" binding:"required"`
	Name         string         `gorm:"type:varchar(255);not null" json:"name" binding:"required"`
	HeightCm     *int           `json:"height_cm,omitempty"`
	WeightKg     *int           `json:"weight_kg,omitempty"`
	Position     string         `gorm:"type:varchar(50);not null;check:position IN ('penyerang', 'gelandang', 'bertahan', 'penjaga gawang')" json:"position" binding:"required,oneof=penyerang gelandang bertahan 'penjaga gawang'"`
	JerseyNumber int            `gorm:"not null" json:"jersey_number" binding:"required,min=1,max=99"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

	// Relasi
	Team  Team   `gorm:"foreignKey:TeamID" json:"team,omitempty"`
	Goals []Goal `gorm:"foreignKey:PlayerID" json:"-"`
}

// TableName menentukan nama tabel untuk model Player
func (Player) TableName() string {
	return "players"
}
