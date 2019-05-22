package entities

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User entity
type User struct {
	gorm.Model
	Username string
	Password string
	Fullname string
	Phone string
	Email string
	LastLogin time.Time
	Gender string
}
