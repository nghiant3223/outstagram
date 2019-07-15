package models

import (
	"github.com/jinzhu/gorm"
)

// PostImage entity
type PostImage struct {
	gorm.Model
	Content       *string
	CommentableID uint
	ViewableID    uint
	ReactableID   uint
	ImageID       uint `gorm:"not null"`
	PostID        uint
	Post          Post
	Image         Image `gorm:"foreignkey:ImageID"`
}
