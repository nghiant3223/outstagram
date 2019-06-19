package storydtos

import "outstagram/server/dtos/dtomodels"

type CreateStoryRequest struct {
}

type CreateStoryResponse struct {
	Stories []dtomodels.Story `json:"stories"`
}
