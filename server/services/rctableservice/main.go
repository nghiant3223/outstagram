package rctableservice

import (
	"outstagram/server/constants"
	postVisibility "outstagram/server/enums/postprivacy"
	"outstagram/server/models"
	"outstagram/server/repos/rctablerepo"
)

type ReactableService struct {
	reactableRepo *rctablerepo.ReactableRepo
}

func New(reactableRepo *rctablerepo.ReactableRepo) *ReactableService {
	return &ReactableService{reactableRepo: reactableRepo}
}

func (s *ReactableService) GetReactorsOrderByQuality(reactableID, userID uint, limit int) []models.User {
	return s.reactableRepo.GetReactorsOrderByQuality(reactableID, userID, limit)
}

func (s *ReactableService) GetReactCount(reactableID uint) int {
	return s.reactableRepo.GetReactCount(reactableID)
}

func (s *ReactableService) GetVisibilityByID(reactableID uint) (postVisibility.Privacy, uint, error) {
	return s.reactableRepo.GetVisibility(reactableID)
}

func (s *ReactableService) GetReactors(reactableID, userID uint, limit int) []models.User {
	return s.reactableRepo.GetReactorsOrderByQuality(reactableID, userID, limit)
}

func (s *ReactableService) GetReactorsFullname(reactableID, userID uint) []string {
	var users []string

	for _, user := range s.GetReactorsOrderByQuality(reactableID, userID, constants.ReactorCount) {
		users = append(users, user.Fullname)
	}

	return users
}

func (s *ReactableService) CheckUserReaction(userID, reactableID uint) bool {
	return s.reactableRepo.CheckUserReaction(userID, reactableID)
}
