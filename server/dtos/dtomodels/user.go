package dtomodels

import "time"

type User struct {
	ID             uint         `json:"id"`
	Username       string       `json:"username"`
	Fullname       string       `json:"fullname"`
	Phone          *string      `json:"phone"`
	Email          string       `json:"email"`
	Gender         bool         `json:"gender"`
	IsMe           bool         `json:"isMe"`
	FollowerCount  int          `json:"followerCount"`
	Followers      []SimpleUser `json:"followers"`
	FollowingCount int          `json:"followingCount"`
	Followings     []SimpleUser `json:"following"`
	Followed       *bool        `json:"followed,omitempty"`
	PostCount      int          `json:"postCount"`
	CreatedAt      time.Time    `json:"createdAt"`
	LastLogin      *time.Time   `json:"lastLogin"`
	LastLogout     *time.Time   `json:"lastLogout"`
}

type SimpleUser struct {
	ID         uint       `json:"id"`
	Username   string     `json:"username"`
	Fullname   string     `json:"fullname"`
	Followed   *bool      `json:"followed"`
	LastLogin  *time.Time `json:"lastLogin"`
	LastLogout *time.Time `json:"lastLogout"`
}
