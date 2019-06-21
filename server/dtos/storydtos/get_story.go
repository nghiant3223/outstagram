package storydtos

import "outstagram/server/dtos/dtomodels"

type GetStoryFeedResponse struct {
	StoryBoards []*dtomodels.StoryBoard `json:"storyBoards"`
}
