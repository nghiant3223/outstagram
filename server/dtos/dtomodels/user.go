package dtomodels

type User struct {
	ID        uint    `json:"id"`
	Username  string  `json:"username"`
	Fullname  string  `json:"fullname"`
	AvatarURL *string `json:"avatarURL"`
	Phone     *string `json:"phone"`
	Email     string  `json:"email"`
	Gender    bool    `json:"gender"`
}

type BasicUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
}
