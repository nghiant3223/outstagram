package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User entity
type User struct {
	gorm.Model
	Username     string  `gorm:"unique;not null;unique_index"`
	Password     string  `gorm:"not null"`
	Fullname     string  `gorm:"not null"`
	Phone        *string `gorm:"unique"`
	Email        string  `gorm:"unique; not null"`
	LastLogin    *time.Time
	Gender       bool
	NotifBoardID uint `gorm:"column:notif_board_id"`
	StoryBoardID uint `gorm:"column:story_board_id"`
}
