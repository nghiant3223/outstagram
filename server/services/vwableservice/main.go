package vwableservice

import (
	"outstagram/server/repos/vwablerepo"
)

type ViewableService struct {
	viewableRepo *vwablerepo.ViewableRepo
}

func New(viewableRepo *vwablerepo.ViewableRepo) *ViewableService {
	return &ViewableService{viewableRepo: viewableRepo}
}

func (s *ViewableService) IncrementView(userID, viewableID uint) error {
	return s.viewableRepo.IncrementView(userID, viewableID)
}

func (s *ViewableService) SaveView(userID, viewableID uint) error {
	return s.viewableRepo.SaveView(userID, viewableID)
}
