package entities

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Notification entity
type Notification struct {
	gorm.Model
	Content string
	CreatedAt string
	State int
	ReadAt time.Time
}