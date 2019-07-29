package roomdtos

import "outstagram/server/dtos/dtomodels"

type CreateRoomRequest struct {
	MemberIDs    []uint   `json:"memberIDs"`
	FirstMessage *Message `json:"1stMessage"`
}

type Message struct {
	Content string `json:"content"`
	Type    int8   `json:"type"`
}

type CreateRoomResponse struct {
	Room dtomodels.Room `json:"room"`
}
