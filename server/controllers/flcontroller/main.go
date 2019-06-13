package flcontroller

import "outstagram/server/services/userservice"

type Controller struct {
	userService *userservice.UserService
}

func New(userService *userservice.UserService) *Controller {
	return &Controller{userService: userService}
}
