package models

import (
	"github.com/jinzhu/gorm"
)

// React entity
type React struct {
	gorm.Model
	ReactableID uint
	UserID      uint
	User        User
}
