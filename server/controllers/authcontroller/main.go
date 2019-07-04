package authcontroller

import (
	"outstagram/server/services/imgservice"
	"outstagram/server/services/notifbservice"
	"outstagram/server/services/storybservice"
	"outstagram/server/services/userservice"
)

type Controller struct {
	userService  *userservice.UserService
	nbService    *notifbservice.NotifBoardService
	sbService    *storybservice.StoryBoardService
	imageService *imgservice.ImageService
}

func New(userService *userservice.UserService, nbService *notifbservice.NotifBoardService, sbService *storybservice.StoryBoardService, imageService *imgservice.ImageService) *Controller {
	return &Controller{userService: userService, nbService: nbService, sbService: sbService, imageService: imageService}
}
