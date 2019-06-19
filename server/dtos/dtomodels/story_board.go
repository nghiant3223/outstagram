package dtomodels

type StoryBoard struct {
	UserID      uint    `json:"userID"`
	Fullname    string  `json:"fullname"`
	AvatarURL   *string  `json:"avatarURL"`
	HasNewStory bool    `json:"hasNewStory"`
	Stories     []Story `json:"stories"`
	StoryCount  int     `json:"storyCount"`
}
