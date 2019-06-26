package usercontroller

import (
	"outstagram/server/services/storybservice"
	"outstagram/server/services/userservice"
)

type Controller struct {
	userService       *userservice.UserService
	storyBoardService *storybservice.StoryBoardService
}

func New(userService *userservice.UserService, storyBoardService *storybservice.StoryBoardService) *Controller {
	return &Controller{userService: userService, storyBoardService: storyBoardService}
}
