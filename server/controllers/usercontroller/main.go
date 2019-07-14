package usercontroller

import (
	"outstagram/server/services/imgservice"
	"outstagram/server/services/postimgservice"
	"outstagram/server/services/postservice"
	"outstagram/server/services/storybservice"
	"outstagram/server/services/userservice"
	"outstagram/server/services/vwableservice"
)

type Controller struct {
	userService       *userservice.UserService
	storyBoardService *storybservice.StoryBoardService
	postService       *postservice.PostService
	imageService      *imgservice.ImageService
	postImageService  *postimgservice.PostImageService
	viewableService   *vwableservice.ViewableService
}

func New(userService *userservice.UserService,
	storyBoardService *storybservice.StoryBoardService,
	postService *postservice.PostService,
	imageService *imgservice.ImageService,
	postImageService *postimgservice.PostImageService,
	viewableService *vwableservice.ViewableService) *Controller {
	return &Controller{
		userService:       userService,
		storyBoardService: storyBoardService,
		postService:       postService,
		imageService:      imageService,
		postImageService:  postImageService,
		viewableService:   viewableService,
	}
}
