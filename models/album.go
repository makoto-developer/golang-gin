package models

import (
	"time"

	"gorm.io/gorm"
)

// Album represents an album entity
type Album struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Title     string         `gorm:"size:255;not null" json:"title"`
	Artist    string         `gorm:"size:255;not null" json:"artist"`
	Price     float64        `gorm:"type:decimal(10,2);not null" json:"price"`
	Tax       float32        `gorm:"type:decimal(4,2);not null;default:0.1" json:"tax"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for Album model
func (Album) TableName() string {
	return "albums"
}
