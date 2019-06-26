package userdtos

import "outstagram/server/dtos/dtomodels"

type GetStoryBoardResponse struct {
	StoryBoard *dtomodels.StoryBoard `json:"storyBoard"`
}
