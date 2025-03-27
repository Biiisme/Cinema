package model

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	ID        uint           `gorm:"primaryKey"`
	UserID    string         `gorm:"not null" json:"UserID"`
	IP        string         `gorm:"not null" json:"IP"`
	Token     string         `gorm:"not null" json:"token"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
