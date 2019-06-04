package models

import (
	"github.com/jinzhu/gorm"
)

// PostImage entity
type PostImage struct {
	gorm.Model
	CommentableID uint
	ViewableID    uint
	ReactableID   uint
	ImageID       uint `gorm:"not null"`
	PostID        uint
	Image         Image `gorm:"foreignkey:ImageID"`
}
