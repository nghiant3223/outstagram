package postcontroller

import (
	"outstagram/server/services/cmtableservice"
	"outstagram/server/services/cmtservice"
	"outstagram/server/services/imgservice"
	"outstagram/server/services/postimgservice"
	"outstagram/server/services/postservice"
	"outstagram/server/services/rctableservice"
)

type Controller struct {
	postService        *postservice.PostService
	imageService       *imgservice.ImageService
	postImageService   *postimgservice.PostImageService
	commentableService *cmtableservice.CommentableService
	commentService     *cmtservice.CommentService
	reactableService   *rctableservice.ReactableService
}

func New(postService *postservice.PostService, imageService *imgservice.ImageService, postImageService *postimgservice.PostImageService, commentableService *cmtableservice.CommentableService, commentService *cmtservice.CommentService, reactableService *rctableservice.ReactableService) *Controller {
	return &Controller{postService: postService, imageService: imageService, postImageService: postImageService, commentableService: commentableService, commentService: commentService, reactableService: reactableService}
}
