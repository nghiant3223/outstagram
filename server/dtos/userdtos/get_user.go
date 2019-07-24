package userdtos

import "outstagram/server/dtos/dtomodels"

type GetUserResponse struct {
	User dtomodels.User `json:"user"`
}
