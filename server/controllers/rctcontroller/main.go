package rctcontroller

import (
	"outstagram/server/services/rctservice"
)

//type Controller struct {
//	reactableService *rctableservice.ReactableService
//	reactService     *rctservice.ReactService
//	postService      *postservice.PostService
//	commentService   *cmtservice.CommentService
//}
//
//func New(reactableService *rctableservice.ReactableService, reactService *rctservice.ReactService, postService *postservice.PostService, commentService *cmtservice.CommentService) *Controller {
//	return &Controller{reactableService: reactableService, reactService: reactService, postService: postService, commentService: commentService}
//}

type Controller struct {
	reactService *rctservice.ReactService
}

func New(reactService *rctservice.ReactService, ) *Controller {
	return &Controller{reactService: reactService}
}
