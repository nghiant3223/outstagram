package roomcontroller

import (
	"outstagram/server/services/roomservice"
)

type Controller struct {
	roomService *roomservice.RoomService
}

func New(roomService *roomservice.RoomService) *Controller {
	return &Controller{
		roomService: roomService,
	}
}
