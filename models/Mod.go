package models

import "gorm.io/gorm"

type Mod struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
