package models

import (
	"github.com/jinzhu/gorm"
	"outstagram/server/enums/postprivacy"
)

// Post entity
type Post struct {
	gorm.Model
	Content       *string
	NumViewed     int                 `gorm:"default:0"`
	Privacy       postPrivacy.Privacy `gorm:"default:0"`
	CommentableID uint
	ReactableID   uint
	ViewableID    uint
	UserID        uint
	User          User
	Images        []PostImage `gorm:"foreignkey:PostID"`
	ImageCount    int         `gorm:"-"`
	Popularity    float32     `gorm:"popularity;default:0	"`
}
