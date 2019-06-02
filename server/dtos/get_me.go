package dtos

import "time"

type GetMeResponse struct {
	Username     string     `json:"username"`
	Fullname     string     `json:"fullname"`
	Email        string     `json:"email"`
	Phone        *string    `json:"phone"`
	LastLogin    *time.Time `json:"lastLogin"`
	Gender       bool       `json:"gender"`
	NumFollower  int        `json:"numFollower"`
	NumFollowing int       `json:"numFollowing"`
	NotifBoardID uint       `json:"notifBoardID"`
	StoryBoardID uint       `json:"storyBoardID"`
}
