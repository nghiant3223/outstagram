package entities

import (
	"github.com/jinzhu/gorm"
)

// Reply entity
type Reply struct {
	gorm.Model
	Content string
}