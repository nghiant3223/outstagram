package entities

import (
	"github.com/jinzhu/gorm"
)

// StoryImage entity
type StoryImage struct {
	gorm.Model
	ImageID uint
}
