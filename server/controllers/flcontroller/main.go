package flcontroller

import (
	"outstagram/server/services/postservice"
	"outstagram/server/services/userservice"
)

type Controller struct {
	userService *userservice.UserService
	postService *postservice.PostService
}

func New(userService *userservice.UserService,
	postService *postservice.PostService) *Controller {
	return &Controller{
		userService: userService,
		postService: postService,
	}
}
