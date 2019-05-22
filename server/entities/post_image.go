package entities

import (
	"github.com/jinzhu/gorm"
)

// PostImage entity
type PostImage struct {
	gorm.Model
	Content string
	CommentableID uint
	ReactableID uint
	ImageID uint
}