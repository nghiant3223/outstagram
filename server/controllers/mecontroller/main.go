package mecontroller

import (
	"outstagram/server/services/imgservice"
	"outstagram/server/services/postservice"
	"outstagram/server/services/storybservice"
	"outstagram/server/services/userservice"
)

type Controller struct {
	userService       *userservice.UserService
	postService       *postservice.PostService
	storyBoardService *storybservice.StoryBoardService
	imageService      *imgservice.ImageService
}

func New(userService *userservice.UserService, postService *postservice.PostService, storyBoardService *storybservice.StoryBoardService, imageService *imgservice.ImageService) *Controller {
	return &Controller{userService: userService, postService: postService, storyBoardService: storyBoardService, imageService: imageService}
}
