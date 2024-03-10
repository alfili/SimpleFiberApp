package models

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	UserID      uint
	Description string `json:"description"`
}
