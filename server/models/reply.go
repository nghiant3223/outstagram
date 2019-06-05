package models

import (
	"github.com/jinzhu/gorm"
)

// Reply entity
type Reply struct {
	gorm.Model
	Content   string `gorm:"not null"`
	CommentID uint
	UserID    uint
}
