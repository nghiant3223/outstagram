package postcontroller

import (
	"outstagram/server/services/imgservice"
	"outstagram/server/services/postimgservice"
	"outstagram/server/services/postservice"
)

type Controller struct {
	postService      *postservice.PostService
	imageService     *imgservice.ImageService
	postImageService *postimgservice.PostImageService
}

func New(postService *postservice.PostService, imageService *imgservice.ImageService, postImageService *postimgservice.PostImageService) *Controller {
	return &Controller{postService: postService, imageService: imageService, postImageService: postImageService}
}
