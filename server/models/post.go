package models

import (
	"github.com/jinzhu/gorm"
	"outstagram/server/enums/postenums"
)

// Post entity
type Post struct {
	gorm.Model
	Content       *string
	NumRead       int                  `gorm:"default:0"`
	Visibility    postenums.Visibility `gorm:"default:0"`
	CommentableID uint
	ReactableID   uint
	ViewableID    uint
	PostImages    []PostImage `gorm:"foreignkey:PostID"`
	UserID        uint
}
