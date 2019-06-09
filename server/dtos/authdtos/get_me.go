package authdtos

import "time"

type GetMeResponse struct {
	Username       string     `json:"username"`
	Fullname       string     `json:"fullname"`
	Email          string     `json:"email"`
	Phone          *string    `json:"phone"`
	LastLogin      *time.Time `json:"lastLogin"`
	Gender         bool       `json:"gender"`
	FollowerCount  int        `json:"numFollower"`
	FollowingCount int        `json:"numFollowing"`
	NotifBoardID   uint       `json:"notifBoardID"`
	StoryBoardID   uint       `json:"storyBoardID"`
}
