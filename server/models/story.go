package models

import (
	"github.com/jinzhu/gorm"
)

// Story entity
type Story struct {
	gorm.Model
	ImageID uint  `gorm:"not null"`
	Image   Image `gorm:"foreignkey:ImageID"`
}
