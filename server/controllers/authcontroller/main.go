package authcontroller

import (
	"outstagram/server/services/nbservice"
	"outstagram/server/services/sbservice"
	"outstagram/server/services/userservice"
)

type Controller struct {
	userService *userservice.UserService
	nbService   *nbservice.NotifBoardService
	sbService   *sbservice.StoryBoardService
}

func New(userService *userservice.UserService, nbService *nbservice.NotifBoardService, sbService *sbservice.StoryBoardService) *Controller {
	return &Controller{userService: userService, nbService: nbService, sbService: sbService}
}
