package userdtos

type GetUserResponse struct {
	ID             uint    `json:"id"`
	IsMe           bool    `json:"isMe"`
	Username       string  `json:"username"`
	Fullname       string  `json:"fullname"`
	FollowerCount  int     `json:"followerCount"`
	FollowingCount int     `json:"followingCount"`
	Followed       *bool   `json:"followed,omitempty"`
}
