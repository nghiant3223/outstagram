package models

import (
	"github.com/jinzhu/gorm"
)

// Message entity
type Message struct {
	gorm.Model
	Content string `gorm:"not null"`
	Type int8 `gorm:"default:0"`
}