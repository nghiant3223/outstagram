package imgcontroller

import (
	"outstagram/server/services/imgservice"
	"outstagram/server/services/userservice"
)

type Controller struct {
	imageService *imgservice.ImageService
	userService  *userservice.UserService
}

func New(imageService *imgservice.ImageService, userService *userservice.UserService) *Controller {
	return &Controller{imageService: imageService, userService: userService}
}
