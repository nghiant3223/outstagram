package entities

import (
	"github.com/jinzhu/gorm"
)

// Comment entity
type Comment struct {
	gorm.Model
	Content string
}