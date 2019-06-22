package dtomodels

type StoryBoard struct {
	UserID      uint    `json:"userID"`
	Fullname    string  `json:"fullname"`
	AvatarURL   *string `json:"avatarURL"`
	IsMy        bool    `json:"isMy"`
	HasNewStory bool    `json:"hasNewStory"`
	Stories     []Story `json:"stories"`
	StoryCount  int     `json:"storyCount"`
}
