package usercontroller

import (
	userservice "outstagram/server/services/userservice"
)

type Controller struct {
	service *userservice.UserService
}

func New(userService *userservice.UserService) *Controller {
	return &Controller{service: userService}
}
