package models

import (
	"github.com/jinzhu/gorm"
)

// Comment entity
type Comment struct {
	gorm.Model
	Content       *string `gorm:"not null"`
	CommentableID uint
	ReactableID   uint
	UserID        uint
	User          User
	Replies       []Reply
	ReplyCount    int `gorm:"-"`
}
