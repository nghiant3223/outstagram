package medtos

import "outstagram/server/dtos/dtomodels"

type GetNewsFeedResponse struct {
	Posts []dtomodels.Post `json:"posts"`
}
