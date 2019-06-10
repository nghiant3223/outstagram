package viewableservice

import (
	"outstagram/server/repos/userrepo"
	"outstagram/server/repos/viewablerepo"
)

type ViewableService struct {
	viewableRepo *viewablerepo.ViewableRepo
}

func New(userRepo *userrepo.UserRepo) *ViewableService {
	return &ViewableService{userRepo: userRepo}
}

func (s *ViewableService) IncrementView(userID, viewableID uint) error {
	return s.viewableRepo.IncrementView(userID, viewableID)
}
