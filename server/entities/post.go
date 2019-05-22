package entities

import (
	"github.com/jinzhu/gorm"
)

// Post entity
type Post struct {
	gorm.Model
	Content string
	CommentableID uint
	ReactableID uint
	NumRead int
}