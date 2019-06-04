package followdtos

type Follower struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatarURL"`
}

type Following struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatarURL"`
}
