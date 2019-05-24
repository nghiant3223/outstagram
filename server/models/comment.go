package models

import (
	"github.com/jinzhu/gorm"
)

// Comment entity
type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
}