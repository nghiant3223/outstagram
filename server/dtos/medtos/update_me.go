package medtos

import "outstagram/server/dtos/dtomodels"

type UpdateMeRequest struct {
}

type UpdateMeResponse struct {
	Me dtomodels.Me `json:"me"`
}
