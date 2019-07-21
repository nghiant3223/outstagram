package medtos

import "outstagram/server/dtos/dtomodels"

type GetNewsFeedResponse struct {
	Posts       []dtomodels.Post `json:"posts"`
	NextSinceID uint            `json:"nextSinceID,omitempty"` // Next post's ID to fetch
}

type GetNewsFeedRequest struct {
	SinceID    uint `form:"since_id"`
	Pagination bool `form:"pagination"`
}
