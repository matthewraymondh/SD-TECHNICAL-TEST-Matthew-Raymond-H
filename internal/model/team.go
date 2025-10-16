package model

import (
	"time"

	"gorm.io/gorm"
)

// Team merepresentasikan tabel teams di database
type Team struct {
	ID                  uint           `gorm:"primaryKey" json:"id"`
	Name                string         `gorm:"type:varchar(255);not null" json:"name" binding:"required"`
	LogoURL             *string        `gorm:"type:varchar(255)" json:"logo_url,omitempty"`
	FoundedYear         *int           `json:"founded_year,omitempty"`
	HeadquartersAddress *string        `gorm:"type:text" json:"headquarters_address,omitempty"`
	HeadquartersCity    *string        `gorm:"type:varchar(100)" json:"headquarters_city,omitempty"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"-"`

	// Relasi
	Players     []Player `gorm:"foreignKey:TeamID" json:"players,omitempty"`
	HomeMatches []Match  `gorm:"foreignKey:HomeTeamID" json:"-"`
	AwayMatches []Match  `gorm:"foreignKey:AwayTeamID" json:"-"`
}

// TableName menentukan nama tabel untuk model Team
func (Team) TableName() string {
	return "teams"
}
