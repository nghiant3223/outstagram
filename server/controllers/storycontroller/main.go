package storycontroller

import (
	"outstagram/server/services/imgservice"
	"outstagram/server/services/storybservice"
	"outstagram/server/services/userservice"
	"outstagram/server/services/vwableservice"
)

type Controller struct {
	imageService      *imgservice.ImageService
	viewableService   *vwableservice.ViewableService
	storyBoardService *storybservice.StoryBoardService
	userService       *userservice.UserService
}

func New(imageService *imgservice.ImageService, viewableService *vwableservice.ViewableService, storyBoardService *storybservice.StoryBoardService, userService *userservice.UserService) *Controller {
	return &Controller{imageService: imageService, viewableService: viewableService, storyBoardService: storyBoardService, userService: userService}
}
