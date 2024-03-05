package models

import (
	"time"

	"gorm.io/gorm"
)

type Mod struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `json:"name" gorm:"unique"`
	Description string         `json:"description"`
}
