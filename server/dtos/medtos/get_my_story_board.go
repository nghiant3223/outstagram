package medtos

import "outstagram/server/dtos/dtomodels"

type GetMyStoryBoard struct {
	Stories []dtomodels.Story `json:"stories"`
}
