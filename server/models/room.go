package models

import (
	"github.com/jinzhu/gorm"
)

// Room entity
type Room struct {
	gorm.Model
	Name string
	Type bool `gorm:"default:false"`
}