package rctableservice

import (
	"outstagram/server/constants"
	postVisibility "outstagram/server/enums/postvisibility"
	"outstagram/server/models"
	"outstagram/server/repos/rctablerepo"
)

type ReactableService struct {
	reactableRepo *rctablerepo.ReactableRepo
}

func New(reactableRepo *rctablerepo.ReactableRepo) *ReactableService {
	return &ReactableService{reactableRepo: reactableRepo}
}

func (s *ReactableService) GetReactsWithLimit(id , limit , offset uint) (*models.Reactable, error) {
	return s.reactableRepo.GetReactsWithLimit(id, limit, offset)
}

func (s *ReactableService) GetReactors(reactableID, userID uint, limit int) []models.User {
	return s.reactableRepo.GetReactors(reactableID, userID, limit)
}

func (s *ReactableService) GetReactCount(reactableID uint) int {
	return s.reactableRepo.GetReactCount(reactableID)
}

func (s *ReactableService) GetVisibilityByID(reactableID uint) (postVisibility.Visibility, uint, error) {
	return s.reactableRepo.GetVisibility(reactableID)
}

func (s *ReactableService) GetReactorsFullname(reactableID, userID uint) []string {
	var users []string

	for _, user := range s.GetReactors(reactableID, userID, constants.ReactorCount) {
		users = append(users, user.Fullname)
	}

	return users
}