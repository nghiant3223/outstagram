package roomdtos

import "outstagram/server/dtos/dtomodels"

type GetRecentRoomResponse struct {
	Rooms []dtomodels.Room `json:"rooms"`
}
