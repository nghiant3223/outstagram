package viewableservice

import (
	"outstagram/server/repos/viewablerepo"
)

type ViewableService struct {
	viewableRepo *viewablerepo.ViewableRepo
}

func New(viewableRepo *viewablerepo.ViewableRepo) *ViewableService {
	return &ViewableService{viewableRepo: viewableRepo}
}

func (s *ViewableService) IncrementView(userID, viewableID uint) error {
	return s.viewableRepo.IncrementView(userID, viewableID)
}
