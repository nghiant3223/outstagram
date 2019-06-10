package rctcontroller

import (
	"outstagram/server/services/rctableservice"
)

type Controller struct {
	reactableService *rctableservice.ReactableService
}

func New(reactableService *rctableservice.ReactableService) *Controller {
	return &Controller{reactableService: reactableService}
}


