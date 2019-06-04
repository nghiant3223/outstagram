package models

import (
	"github.com/jinzhu/gorm"
)

// StoryImage entity
type StoryImage struct {
	gorm.Model
	ImageID uint  `gorm:"not null"`
	Image   Image `gorm:"foreignkey:ImageID"`
}
