package dtomodels

import "time"

type Me struct {
	ID           uint       `json:"id"`
	Username     string     `json:"username"`
	Fullname     string     `json:"fullname"`
	Phone        *string    `json:"phone"`
	Email        string     `json:"email"`
	Gender       bool       `json:"gender"`
	LastLogin    *time.Time `json:"lastLogin"`
	LastLogout   *time.Time `json:"lastLogout"`
	CoverImageID uint       `json:"coverImageID"`
	CreatedAt    time.Time  `json:"createdAt"`
}
