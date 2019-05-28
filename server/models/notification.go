package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Notification entity
type Notification struct {
	gorm.Model
	Content   string `gorm:"not null"`
	State     int `gorm:"default:0"`
	ReadAt    time.Time
}