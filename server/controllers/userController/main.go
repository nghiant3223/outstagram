package usercontroller

import (
	"outstagram/server/services"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{service: userService}
}
