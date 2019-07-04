package dtomodels

type StoryBoard struct {
	StoryBoardID uint    `json:"storyBoardID"`
	UserID       uint    `json:"userID"`
	Fullname     string  `json:"fullname"`
	IsMy         bool    `json:"isMy"`
	HasNewStory  bool    `json:"hasNewStory"`
	Stories      []Story `json:"stories"`
	StoryCount   int     `json:"storyCount"`
}
