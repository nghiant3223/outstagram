package storydtos

import "outstagram/server/dtos/dtomodels"

type CreateStoryRequest struct {
	ImageURLs []string `form:"imageURLs"`
}

type CreateStoryResponse struct {
	Stories []dtomodels.Story `json:"stories"`
}
