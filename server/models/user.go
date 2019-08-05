package models

import (
	"log"
	"outstagram/server/dtos/dtomodels"
	"regexp"
	"time"

	"github.com/jinzhu/gorm"
)

// User entity
type User struct {
	gorm.Model
	Username      string  `gorm:"unique;not null;unique_index"`
	Password      string  `gorm:"not null"`
	Fullname      string  `gorm:"not null"`
	Phone         *string `gorm:"unique"`
	Email         string  `gorm:"unique; not null"`
	LastLogin     *time.Time
	LastLogout    *time.Time
	Gender        bool
	Rooms         []*Room    `gorm:"many2many:joins"`
	NotifBoard    NotifBoard `gorm:"association_autoupdate:false"`
	StoryBoard    StoryBoard `gorm:"association_autoupdate:false"`
	AvatarImageID uint
	CoverImageID  uint
}

func (u *User) ToDTO() dtomodels.User {
	return dtomodels.User{
		ID:           u.ID,
		Fullname:     u.Fullname,
		Username:     u.Username,
		Gender:       u.Gender,
		Phone:        u.Phone,
		Email:        u.Email,
		LastLogout:   u.LastLogout,
		LastLogin:    u.LastLogin,
		CoverImageID: u.CoverImageID,
		CreatedAt:    u.CreatedAt,
	}
}

func (u *User) ToMeDTO() dtomodels.Me {
	return dtomodels.Me{
		ID:           u.ID,
		Fullname:     u.Fullname,
		Username:     u.Username,
		Gender:       u.Gender,
		Phone:        u.Phone,
		Email:        u.Email,
		LastLogout:   u.LastLogout,
		LastLogin:    u.LastLogin,
		CoverImageID: u.CoverImageID,
		CreatedAt:    u.CreatedAt,
	}
}

func (u *User) ToSimpleDTO() dtomodels.SimpleUser {
	return dtomodels.SimpleUser{
		ID:         u.ID,
		Fullname:   u.Fullname,
		Username:   u.Username,
		LastLogout: u.LastLogout,
		LastLogin:  u.LastLogin,
		CreatedAt:  u.CreatedAt,
	}
}

func (u *User) IsValid() (bool, string) {
	usernameOK, err := regexp.MatchString(`^[a-zA-Z][a-zA-Z0-9._]{5,}$`, u.Username)
	if err != nil {
		log.Print("Invalid regex", err.Error())
		return false, ""
	}

	if !usernameOK {
		return false, "Invalid username. Password must be greater than 5 characters in length and contains at least a letter"
	}

	passwordOK := len(u.Password) >= 5
	if !passwordOK {
		return false, "Invalid password. Password must be greater than 5 characters in length"
	}

	return true, ""
}
